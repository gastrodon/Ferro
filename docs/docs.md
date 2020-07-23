<details>
<summary>POST /</summary>
Upload some file. POST body should be multipart data containing the mime data to upload

__headers__

|name|value|required|
| - | - | - |
|Authorization|Auth header that was created by you when configuring this server. Servers should only have one application talking to them at a time, though this may be from a number of nodes with different auth headers|True|

__responses__

- 200 - File uploaded
```JSON
{
    "id": "UUIDv4 that points to this file"
}
```


</details>


<details>
<summary>GET /:id/</summary>
Get some stored file by its id

__headers__

|name|value|required|
| - | - | - |
|Authorization|Auth header that was created by you when configuring this server. Servers should only have one application talking to them at a time, though this may be from a number of nodes with different auth headers|True|

__responses__

- 200 - Requested file
- 404 - No file exists
```JSON
{
    "error": "not_found"
}
```


</details>
<details>
<summary>POST /:id/</summary>
Upload some file with a predefined name. Will repalce any file with this name already

__headers__

|name|value|required|
| - | - | - |
|Authorization|Auth header that was created by you when configuring this server. Servers should only have one application talking to them at a time, though this may be from a number of nodes with different auth headers|True|

__responses__

- 200 - File uploaded
```JSON
{
    "id": "UUIDv4 that points to this file"
}
```


</details>
<details>
<summary>DELETE /:id/</summary>
Delete some stored file by its id

__headers__

|name|value|required|
| - | - | - |
|Authorization|Auth header that was created by you when configuring this server. Servers should only have one application talking to them at a time, though this may be from a number of nodes with different auth headers|True|

__responses__

- 204 - File deleted

</details>
