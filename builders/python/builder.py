from flask import (Flask, request, abort)
import logging
from utils import (install_package, uninstall_package,
                   get_package_info, get_package_size,
                   get_package_location, create_error,
                   make_package_path, remove_package_path,
                   get_install_path, print_flush)

app = Flask(__name__)


@ app.route('/size')
def size():
    package = request.form['package']
    package = package.lower()
    try:
        package_path = get_install_path(package)
        make_package_path(package_path)
        install_package(package_path, package)
        package_info = get_package_info(package_path, package)
        package_location = get_package_location(package_info)
        package_size = get_package_size(package_location, package)
        uninstall_package(package_path, package)
        remove_package_path(package_path)
    except Exception as e:
        return create_error(e)

    return {
        "size": package_size,
        'error': None
    }


if __name__ == '__main__':
    app.run()
