# Clicker Game
화면을 클릭하여 게임내 재화를 벌고 성장하는 게임입니다. 

### Architecture
![alt text](./image/arch.png)

### Build
```bash
git clone https://github.com/txuna/clicker-game.git
cd clicker-game/server
make kind-create

make deploy-base-all
make deploy-app-all
```

### DEMO
![alt text](./image/k9s.png)

![alt text](./image/image.png)

![alt text](./image/image-1.png)

### Cleanup 
```bash
make remove-app-all
make remove-base-all
```