import * as pulumi from "@pulumi/pulumi";
import * as k3os from "@spigell/pulumi-k3os";

let serverName = "k3os-server01"
let agentName = "k3os-agent01"

let config = new pulumi.Config();
let serverIP = config.require(`${serverName}IP`);
let agentIP = config.require(`${agentName}IP`);

const server = new k3os.Node(serverName, {
    connection: {
        addr: `${serverIP}:22`,
        user: "rancher",
        key: `./.vagrant/machines/${serverName}/virtualbox/private_key`
    },
    nodeConfiguration: {
        hostname: "server01",
        k3OS: {
            password: "1234"
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
        hostname: "agent01"
    }
});