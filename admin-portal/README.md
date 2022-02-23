python3 -m pip install Django

brew install mysql-client
export PATH="/usr/local/opt/mysql-client/bin:$PATH"
python3 -m pip install mysqlclient

cd /Users/huanglh/go/src/sg-stay-safe.org/site-management/project/
python3 manage.py runserver

deploy to local:
http://127.0.0.1:8000/admin/

cd /Users/huanglh/go/src/sg-stay-safe.org/admin-portal/project
python3 manage.py makemigrations
python3 manage.py migrate