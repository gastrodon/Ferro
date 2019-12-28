<details>
<summary>POST /</summary>
Upload some file. POST body should be multipart data containing the mime data to upload

__headers__

|name|value|required|
| - | - | - |
|Authorization|Auth header that was created by you when configuring this server. Servers should only have one application talking to them at a time, though this may be from a number of nodes with different auth headers|True|

__responses__

- 200 - File uploaded
File uploaded

The file sent to the server was recieved and stored, and a token was dispatched to the requester

```JSON
{
    "id": "UUIDv4 that points to this file"
}
```


</details>


<details>
<summary>GET /:id/</summary>
Get some stored file by its id

__query strings__

|name|value|type|default|
| - | - | - | - |
|cropx|Pixel crop on both ends of the x axis|int|0|
|cropy|Pixel crop on both ends of the y axis|int|0|
|scale|Image scale percent. Thumbnails use 30(?)|int > 0|100|
|scalex|Pixel scale on x axis|int| |
|scaley|Pixel scale on y axis|int| |

__headers__

|name|value|required|
| - | - | - |
|Authorization|Auth header that was created by you when configuring this server. Servers should only have one application talking to them at a time, though this may be from a number of nodes with different auth headers|True|

__path arguments__

|name|value|required|
| - | - | - |
|id|UUIDv4 of the object being queried|True|

__responses__

- 200 - Requested file
Requested file

The binary of the requested file


</details>
<details>
<summary>DELETE /:id/</summary>
Delete some stored file by its id

__headers__

|name|value|required|
| - | - | - |
|Authorization|Auth header that was created by you when configuring this server. Servers should only have one application talking to them at a time, though this may be from a number of nodes with different auth headers|True|

__path arguments__

|name|value|required|
| - | - | - |
|id|UUIDv4 of the object being queried|True|

__responses__

- 204 - File deleted
File deleted

The queried file was destroyed on the server


</details>


<details>
<summary>GET /:id/md5</summary>
Get the md5 hash of some file by its id

__headers__

|name|value|required|
| - | - | - |
|Authorization|Auth header that was created by you when configuring this server. Servers should only have one application talking to them at a time, though this may be from a number of nodes with different auth headers|True|

__path arguments__

|name|value|required|
| - | - | - |
|id|UUIDv4 of the object being queried|True|

__responses__

- 200 - Content md5
Content md5

The md5 digest of the queried file

```JSON
{
    "md5": "md5 of this file"
}
```


</details>


<details>
<summary>GET /:id/thumb</summary>
Get the thumbnail of this image. This is a 512x512 version of this image

__headers__

|name|value|required|
| - | - | - |
|Authorization|Auth header that was created by you when configuring this server. Servers should only have one application talking to them at a time, though this may be from a number of nodes with different auth headers|True|

__path arguments__

|name|value|required|
| - | - | - |
|id|UUIDv4 of the object being queried|True|

__responses__

- 200 - Requested file
Requested file

The binary of the requested file


</details>
