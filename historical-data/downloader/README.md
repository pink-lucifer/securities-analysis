# Downloader for securities historical data




## Note for development
1- while utilizing github.com/go-sql-driver/mysql for mysql8, met connection issue "this authentication plugin is not supported"
    The workaround is to change authentication plugin for the user.
    However, the root user is created during a previous run of the server, where the default plugin was caching_sha2_password.
    After changing the authentication plugin in the user, you need to set the password again. With ALTER USER you can do both things at once.
    ```
    alter user root@'localhost' identified with mysql_native_password by 'my-secret-pw';
    ```
    for user other than root, the alter statement should work like a charm.
    detail please refer to [issues-785](https://github.com/go-sql-driver/mysql/issues/785)
2- while utilizing github.com/go-sql-driver/mysql for mysql8, met connection issue "Cannot connect to MySQL: this user requires mysql native password authentication."
    It has been suggested in other issue [#625](https://github.com/go-sql-driver/mysql/issues/625)
    to do this to make native password accounts to work:
    Note that you will need to explicitly set "allowNativePasswords=True" in the DSN.