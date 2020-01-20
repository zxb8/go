# 项目依赖的包
```she
go get github.com/joho/godotenv
go get -u github.com/jinzhu/gorm
go get -u github.com/go-sql-driver/mysql
go get -u golang.org/x/crypto/bcrypt

http://gorm.io/
http://gorm.io/docs/connecting_to_the_database.html
```





# 运行项目

```
 go run build main/main.go
```



```shell

GOROOT=C:\Go #gosetup
GOPATH=D:\www\go-learn #gosetup
C:\Go\bin\go.exe build -o C:\Users\31429\AppData\Local\Temp\___1go_build_blogApp_main.exe blogApp/main #gosetup
"C:\Program Files\JetBrains\GoLand 2018.3.4\bin\runnerw64.exe" C:\Users\31429\AppData\Local\Temp\___1go_build_blogApp_main.exe #gosetup

(D:/www/go-learn/src/blogApp/auto/load.go:23) 
[2020-01-20 14:40:57]  [23.94ms]  CREATE TABLE `users` (`id` int unsigned AUTO_INCREMENT,`nickname` varchar(20) NOT NULL UNIQUE,`email` varchar(50) NOT NULL UNIQUE,`password` varchar(60) NOT NULL,`create_at` DATETIME NULL DEFAULT current_timestamp(),`update_at` DATETIME NULL DEFAULT current_timestamp() , PRIMARY KEY (`id`))  
[0 rows affected or returned ] 

(D:/www/go-learn/src/blogApp/auto/load.go:29) 
[2020-01-20 14:40:57]  [140.22ms]  INSERT INTO `users` (`nickname`,`email`,`password`) VALUES ('Jhon Doe','john@163.com','123456')  
[1 rows affected or returned ] 

(D:/www/go-learn/src/blogApp/auto/load.go:29) 
[2020-01-20 14:40:57]  [2.99ms]  SELECT `create_at`, `update_at` FROM `users`  WHERE (id = 1)  
[1 rows affected or returned ] 
{
 "id": 1,
 "nickname": "Jhon Doe",
 "emmail": "john@163.com",
 "password": "123456",
 "created_at": "2020-01-20T14:40:58+08:00",
 "update_at": "2020-01-20T14:40:58+08:00"
}

(D:/www/go-learn/src/blogApp/auto/load.go:29) 
[2020-01-20 14:40:57]  [4.97ms]  INSERT INTO `users` (`nickname`,`email`,`password`) VALUES ('jack','jack@163.com','123456')  
[1 rows affected or returned ] 

(D:/www/go-learn/src/blogApp/auto/load.go:29) 
[2020-01-20 14:40:57]  [6.97ms]  SELECT `create_at`, `update_at` FROM `users`  WHERE (id = 2)  
[1 rows affected or returned ] 
{
 "id": 2,
 "nickname": "jack",
 "emmail": "jack@163.com",
 "password": "123456",
 "created_at": "2020-01-20T14:40:58+08:00",
 "update_at": "2020-01-20T14:40:58+08:00"
}

(D:/www/go-learn/src/blogApp/auto/load.go:29) 
[2020-01-20 14:40:57]  [1.01ms]  INSERT INTO `users` (`nickname`,`email`,`password`) VALUES ('lucy','lucy@163.com','123456')  
[1 rows affected or returned ] 

(D:/www/go-learn/src/blogApp/auto/load.go:29) 
[2020-01-20 14:40:57]  [0.97ms]  SELECT `create_at`, `update_at` FROM `users`  WHERE (id = 3)  
[1 rows affected or returned ] 
{
 "id": 3,
 "nickname": "lucy",
 "emmail": "lucy@163.com",
 "password": "123456",
 "created_at": "2020-01-20T14:40:58+08:00",
 "update_at": "2020-01-20T14:40:58+08:00"
}

	Listening [::]:9000


```

# 添加用户

![创建用户](C:\Users\31429\Desktop\Image 1.png)