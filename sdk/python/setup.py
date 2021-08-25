# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import errno
from setuptools import setup, find_packages
from setuptools.command.install import install
from subprocess import check_call


class InstallPluginCommand(install):
    def run(self):
        install.run(self)
        try:
            check_call(['pulumi', 'plugin', 'install', 'resource', 'k3os', '${PLUGIN_VERSION}', '--server', 'https://github.com/spigell/pulumi-k3os/releases/download/v0.0.2'])
        except OSError as error:
            if error.errno == errno.ENOENT:
                print("""
                There was an error installing the k3os resource provider plugin.
                It looks like `pulumi` is not installed on your system.
                Please visit https://pulumi.com/ to install the Pulumi CLI.
                You may try manually installing the plugin by running
                `pulumi plugin install resource k3os ${PLUGIN_VERSION}`
                """)
            else:
                raise


def readme():
    with open('README.md', encoding='utf-8') as f:
        return f.read()


setup(name='pulumi_k3os',
      version='${VERSION}',
      long_description=readme(),
      long_description_content_type='text/markdown',
      cmdclass={
          'install': InstallPluginCommand,
      },
      keywords='pulumi k3os',
      packages=find_packages(),
      package_data={
          'pulumi_k3os': [
              'py.typed',
          ]
      },
      install_requires=[
          'parver>=0.2.1',
          'pulumi',
          'semver>=2.8.1'
      ],
      zip_safe=False)
