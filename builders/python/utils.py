import subprocess


def run_command(command):
    process = subprocess.Popen(command.split(), stdout=subprocess.PIPE)
    output, error = process.communicate()
    if (error):
        raise Exception(error)
    print(output)
    return output


def install_package(package_name):
    return run_command('pip install {}'.format(package_name))


def uninstall_package(package_name):
    return run_command('pip uninstall {} -y'.format(package_name))


def get_package_info(package):
    return run_command('pip show {}'.format(package))


def get_package_size(package_location, package_name):
    output = run_command(
        'du -sh {}/{}'.format(package_location, package_name))
    return output.decode(encoding="utf-8").split('\t')[0]


def get_package_location(package_info):
    package_info_str = package_info.decode(encoding="utf-8")
    package_location = package_info_str.split("\n")[7][9:]
    return package_location


def create_error(error):
    return {
        'size': None,
        'error': error.args[0]
    }


def print_flush(stmt):
    print(stmt, flush=True)
