# API Document
클라이언트와 서버간의 API 명세서입니다. 

### 로그인
```json 
POST /login 
{
    "user_id": "userid", 
    "password": "password"
}
```
```json 
{
    "error_code" : error_code, 
    "token" : "token string"
}
```

### 돈 채굴
```json 
POST /mining
{
    "token" : "token string"
}
```
```json 
{
    "error_code" : error_code, 
    "player_info" : {
        "coin" : coin,
        "level" : level
    }
}
```

### 유저 정보
```json 
POST /user
{
    "token": "token string"
}
```

```json
{
    "error_code": error_code, 
    "player_info" : {
        "coin": coin,
        "level": level
    }
}
```