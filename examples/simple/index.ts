import * as pulumi from "@pulumi/pulumi";
import * as k3os from "@spigell/pulumi-k3os";
import * as fs from 'fs';

let serverName = "k3os-server01"
let agentName = "k3os-agent01"

let config = new pulumi.Config();
let serverIP = config.require(`${serverName}IP`);
let agentIP = config.require(`${agentName}IP`);

const sshKey = fs.readFileSync('ssh/key.pub', 'utf8');

const server = new k3os.Node(serverName, {
    connection: {
        addr: `${serverIP}:22`,
        user: "rancher",
        key: `./.vagrant/machines/${serverName}/virtualbox/private_key`
    },
    nodeConfiguration: {
        hostname: serverName,
        writeFiles: [
            { path: "/home/rancher/file.txt", content: "Hello" }
        ],
        // Ssh keys are appended instead of overwriting
        sshAuthorizedKeys: [ sshKey ],
        initCmd: [
            "echo This is initCmd"
        ],
        bootCmd: [
            "echo This is bootCmd"
        ],

        runCmd: [
            // Duplicate rancher as vagrant user to let ssh the system with vagrant login
            "sudo sed -e '/^rancher/p' -e 's/^rancher/vagrant/' -i /etc/passwd"
        ],
        k3OS: {
            k3sArgs: [
                "server",
                "--disable-cloud-controller"
            ],
            modules: ["wireguard"],
            sysctls: {
                "kernel.kptr_restrict": "1"
            },
            ntpServers: [ "0.us.pool.ntp.org", "1.us.pool.ntp.org" ],
            token: "token",
            password: "123"
        }
    }
});

const agent = new k3os.Node(agentName, {
    connection: {
        addr: `${agentIP}:22`,
        user: "rancher",
        key: `./.vagrant/machines/${agentName}/virtualbox/private_key`
    },
    nodeConfiguration: {
        hostname: agentName,
        bootCmd: [
            "echo This is bootCmd!"
        ],
        k3OS: {
            token: "token",
            serverUrl: `https://${serverIP}:6443`
        }
    }
});