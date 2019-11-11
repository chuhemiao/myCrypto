package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"os"
)

/*   非对称加密 start  */
//
//// 与对称加密得区别   加密使用得私匙不同
////密匙  是一个密钥对
//// 公钥  可以公开得
//// 私钥  需要保护 丢失后果很严重
//// 加解密都是对应关系，公加私解  私加公解
//
//// 对称加密分发苦难你  对称加密加密数据效率高  对称加密安全级别不如非加
//
//// RSA 加密   密文=  明文 E  mod N    明文得E次方除以N的余数
//// RSA 解密  明文 = 密文 D mod N      密文得D次方除以N的余数
//
//// X.509 通用的证书格式  使用ASN1进行描述数据结构，并使用ASN.1语法进行编码
//// 证书格式   X.509  DER编码  后缀为：.der,.cer,.crt
////  X.509 base64编码（PEM格式），一般后缀为：.pem,.cer,.crt
//// PKCS15个标准  public-key Cryptography Standards  公钥密码标准
//
//
///*   非对称加密 end   */

// 生成公匙和私匙

func RsaGenkey(bits int) error{

	priveKey,err := rsa.GenerateKey(rand.Reader,bits)
	if err != nil{
		panic(err)
	}
	priveStream := x509.MarshalPKCS1PrivateKey(priveKey)
	block := pem.Block{
		Type:    "RSA Pri Key",
		Bytes:   priveStream,
	}
	priFile,err := os.Create("private.pem")
	if err != nil{
		return err
	}
	defer priFile.Close()

	err = pem.Encode(priFile,&block)
	if err != nil{
		return err
	}
	pubKey := priveKey.PublicKey
	// 参数为指针    marshalPublicKey 判断当前的type
	pubStream,err  := x509.MarshalPKIXPublicKey(&pubKey)
	if err != nil{
		return err
	}
	block  = pem.Block{
		Type:    "Pub key",
		Headers: nil,
		Bytes:   pubStream,
	}

	pubFile ,err := os.Create("public.pem")

	if err != nil {
		return err
	}
	defer pubFile.Close()

	// 参数为指针  取值
	err = pem.Encode(pubFile,&block)
	if err != nil {
		return err
	}

	return nil
}

// 公钥加密

func RsaPublicEncrypt(src []byte,pathName string) ([]byte,error){
	// 读公钥文件
	msg := []byte("")
	file, err := os.Open(pathName)
	if err != nil {
		return msg,err
	}
	// 文件属性和文件大小
	info,err := file.Stat()
	if err != nil {
		return msg,err
	}
	recvBuf := make([]byte,info.Size())
	file.Read(recvBuf)
	// 解码
	block,_ := pem.Decode(recvBuf)
	pubInter,err := x509.ParsePKIXPublicKey(block.Bytes)
	if err !=nil{
		return msg,err
	}

	pubkey := pubInter.(*rsa.PublicKey)
	msg,err = rsa.EncryptPKCS1v15(rand.Reader,pubkey,src)
	if err != nil{
		return  msg,err
	}

	return msg,nil

}
// 使用私钥解密

func RsaPrivateDecrypt(src []byte,pathName string) ([]byte,error){
	// 打开私钥文件
	msg := []byte("")
	file,err := os.Open(pathName)
	if err !=nil{
		return msg,err
	}
	// 文件属性和文件大小
	info,err := file.Stat()
	if err != nil {
		return msg,err
	}
	recvBuf := make([]byte,info.Size())
	file.Read(recvBuf)
	// 解码
	block,_ := pem.Decode(recvBuf)
	privateKey,err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil{
		return msg,err
	}
	msg,err = rsa.DecryptPKCS1v15(rand.Reader,privateKey,src)
	if err !=nil{
		return msg,err
	}
	return msg,nil
}
