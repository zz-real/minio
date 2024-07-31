# minio
存储桶列表
get /bucket/list

文件列表
get /bucket/objs/:bucket

文件上传
post /file/get/:bucket  file

下载
post http://localhost:8888/file/get/test 
{
"name": "2.gif",
"path": "F:\\code\\go\\src\\minio"
}

fsnotify 监听目标文件夹的变动，有CREATED就上传