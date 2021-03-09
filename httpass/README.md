# httpass

httpass is an autenthication apps that used to create, remove and get httpasswd file. So, with this app you don't need to add manual to your httpasswd.


#### Create New User

```sh
url: http://localhost:8005
method : POST
Body:

{
    {
        "username":"admin",
        "password":"admin",
        "file_location":"file/demo.htpasswd"
    }
}

```

#### Remove User

```sh
url: http://localhost:8005
method : POST
Body:

{
    {
        "username":"admin",        
        "file_location":"file/demo.htpasswd"
    }
}

```

#### Get User

```sh
url: http://localhost:8005
method : POST
Body:

{
    {
        "file_location":"file/demo.htpasswd"
    }
}

```

