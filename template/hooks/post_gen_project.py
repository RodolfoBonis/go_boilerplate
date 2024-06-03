import os
import shutil

dir_name = 'template'

if os.path.isdir(dir_name):
    for filename in os.listdir(dir_name):
        full_path = os.path.join(dir_name, filename)

        shutil.move(full_path, '.')

    os.rmdir(dir_name)
else:
    print(f'Diretório {dir_name} não encontrado')