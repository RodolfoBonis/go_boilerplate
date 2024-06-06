import os

application_is_service = '{{ cookiecutter.application_is_service }}'

if application_is_service.lower() == 'y':
    os.remove('{{cookiecutter.package_name}}/core/middlewares/auth_middleware.go')
    os.rename('{{cookiecutter.package_name}}/core/middlewares/api_key_auth_middleware.go', '{{cookiecutter.package_name}}/core/middlewares/auth_middleware.go')
else:
    os.remove('{{cookiecutter.package_name}}/core/middlewares/api_key_auth_middleware.go')
