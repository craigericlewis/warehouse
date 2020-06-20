from flask import (Flask, request, abort)
import logging
from utils import (install_package, uninstall_package,
                   get_package_info, get_package_size, get_package_location, create_error)

app = Flask(__name__)


@ app.route('/size')
def size():
    package = request.form['package']
    package = package.lower()
    try:
        install_package(package)
        package_info = get_package_info(package)
        package_location = get_package_location(package_info)
        package_size = get_package_size(package_location, package)
        uninstall_package(package)
    except Exception as e:
        return create_error(e)

    return {
        "size": package_size,
        'error': None
    }


if __name__ == '__main__':
    app.run()
