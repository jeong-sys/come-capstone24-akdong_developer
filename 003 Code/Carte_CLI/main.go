package main

import "carte_cli/cmd"

func main() {
    cmd.Execute()
}

// [권한 부여]
// mkdir /Carte/images

// whoami
// sudo chown -R yj /Carte/images

// =============================여기에 Carte 코드 실행==============================

// terminal실행
// 1) 실행 파일을 시스템 PATH에 추가
// mkdir -p ~/bin
// cp Carte ~/bin/
// echo 'export PATH=$PATH:~/bin' >> ~/.bashrc
// source ~/.bashrc


// 2) '~bin'디렉토리가 시스템 PATH에 포함되어있는가 확인하고 포함되지 않은경우 '.bashrc', '.zshrc', '.profile'에 추가
// echo 'export PATH=$PATH:~/bin' >> ~/.bashrc
// source ~/.bashrc


// 2_2) 'zsh'사용하는 경우 '.zshrc'에 파일 추가
// echo 'export PATH=$PATH:~/bin' >> ~/.zshrc
// source ~/.zshrc
