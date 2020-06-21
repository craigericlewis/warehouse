import subprocess
import uuid
import shutil
import os


def run_command(command, dir="./"):
    process = subprocess.Popen(
        command.split(), stdout=subprocess.PIPE, cwd=dir)
    output, error = process.communicate()
    if (error):
        raise Exception(error)
    return output


def get_install_path(package_name):
    return 'packages/{}-{}'.format(package_name, str(uuid.uuid4()))


def make_package_path(package_path):
    os.mkdir(package_path)


def remove_package_path(package_path):
    shutil.rmtree(package_path)


def install_package(package_path, package_name):
    run_command('pipenv install {}'.format(package_name), package_path)


def uninstall_package(package_path, package_name):
    run_command('pipenv uninstall {} -y'.format(package_name), package_path)


def get_package_info(package_path, package):
    package_info = run_command(
        'pipenv run pip3 show {}'.format(package), package_path)
    return package_info


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
        'error': str(error)
    }


def print_flush(stmt):
    print(stmt, flush=True)
