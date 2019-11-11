package main

import "fmt"

func main(){

	desTest()
	tripleDesTest()
	aesTest()
	rsaTest()
	hashMd5Str()

}

// des测试

func desTest(){

	fmt.Println("=============des 加密=======")
	src := []byte("红红火火恍恍惚惚")
	key := []byte("12345678")
	str := encryptDES(src,key)

	str1 := decryptDES(str,key)

	fmt.Println("解密之后的数据:"+string(str1))
}

// 3des测试

func tripleDesTest(){

	fmt.Println("=========  3des 加密=======")
	src := []byte("红红火火恍恍惚惚,我事谁啊，我是3DES")
	key := []byte("87654321abcdefgh12345678")
	str := encrypt3DES(src,key)

	str1 := decrypt3DES(str,key)

	fmt.Println("解密之后的数据:"+string(str1))
}
// aes测试
func aesTest(){

	fmt.Println("=========  aes 加密=======")
	src := []byte("红红火火恍恍惚惚,我事谁啊，我是AES")
	// 密钥长度为16
	key := []byte("87654321abcdefgh")
	str := encryptAES(src,key)

	str1 := decryptAES(str,key)

	fmt.Println("解密之后的数据:"+string(str1))
}

// rsa 测试

func rsaTest(){
	err := RsaGenkey(4096)
	if err !=nil{
		panic(err)
	}
	// 加密
	src := []byte("红红火火恍恍惚惚,我事谁啊，我是RSA")
	data,err := RsaPublicEncrypt(src,"public.pem")

	data,err = RsaPrivateDecrypt(data,"private.pem")

	if err !=nil{
		panic(err)
	}
	fmt.Println("非对称加密："+string(data))
	// nil为没有错误  看文件是否已经生成
	fmt.Println("错误信息：",err)
}

// hash 测试 md5

func hashMd5Str(){
	fmt.Println("=========  md5 加密=======")
	data  := []byte("红红火火恍恍惚惚,我事谁啊，我是MD5")
	fmt.Println(ShaMd5Str(data))
	fmt.Println(ShaMd5Str2(data))
}

