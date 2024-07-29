OpenSSL官方命令手册：[OpenSSL commands - OpenSSL Documentation](https://docs.openssl.org/master/man1/)

参考教程：

操作：[OpenSSL的基本使用教程(一）_openssl.exe使用教程-CSDN博客](https://blog.csdn.net/fsq0827/article/details/99710052)

操作：[Linux和Shell回炉复习系列文章总目录 - 骏马金龙 - 博客园 (cnblogs.com)](https://www.cnblogs.com/f-ck-need-u/p/7048359.html#auto_id_5)

网站应用：[最新OpenSSL简明教程_openssl使用教程-CSDN博客](https://blog.csdn.net/xxxlllbbb/article/details/107891790)

项目应用：[spring-boot-sign: SpringBoot请求响应加签、验签 (gitee.com)](https://gitee.com/hweiyu/spring-boot-sign)

[最详细的SpringBoot实现接口校验签名调用_springboot接口签名验证-CSDN博客](https://blog.csdn.net/qq_43290318/article/details/131516099)

## 概念

### 公钥和私钥

- **公钥**：可以公开分享，用于加密数据或验证数字签名。
- **私钥**：必须保密，用于解密数据或创建数字签名。
- 关系：公钥和私钥是成对出现的，使用公钥加密的数据只能由对应的私钥解密，反之亦然。

一个完整的RSA私钥包含了以下几个关键参数：

1. 模数（Modulus，N）
2. 公钥指数（Public Exponent，e）
3. 私钥指数（Private Exponent，d）
4. 素数p（Prime 1，p）
5. 素数q（Prime 2，q）
6. 指数dp（Exponent 1，dp）
7. 指数dq（Exponent 2，dq）
8. 系数（Coefficient，qinv）

### 数字签名

- 数字签名使用私钥创建，确保数据的完整性和来源的真实性。
- 验证签名使用公钥，可以确认数据是否被篡改，并且确认数据确实来自特定的私钥持有者。

### 数字证书

- 数字证书由受信任的证书颁发机构（CA）签发，包含公钥及其所有者的身份信息。
- 证书用于验证公钥的真实性，确保其属于声明中的持有者。

### 对称加密和非对称加密

- **对称加密**：使用相同的密钥进行加密和解密，例如 AES、DES。
- **非对称加密**：使用一对公钥和私钥进行加密和解密，例如 RSA。
- 对称加密速度快，适合大数据量；非对称加密安全性高，适合小数据量。

### 证书链

- 证书链由一系列证书组成，确保最终用户证书的可信性。
- 包括根证书（CA）、中间证书和最终用户证书。
- 验证证书链时，必须从根证书到最终用户证书逐级验证。

### 常见文件后缀

**`.pem`**： PEM（Privacy Enhanced Mail）格式文件的标准后缀，通常用于存储证书、私钥、公钥和证书链。文件内容通常以 Base64
编码并用 "-----BEGIN CERTIFICATE-----" 等标头和尾标记。`cert.pem`（证书），`private.pem`（私钥），`public.pem`（公钥）。

**`.key`**：通常用于表示私钥或公钥文件。这个后缀比较通用，可以表示任何类型的加密密钥。 `private.key`（私钥），`public.key`（公钥）。

**`.pub`**:通常用于表示公钥文件。这个后缀明确指示文件包含公钥。`id_rsa.pub`（RSA 公钥），`id_ecdsa.pub`（ECDSA 公钥）。

**`.pri`**：通常用于表示私钥文件。虽然不是标准后缀，但一些系统或工具可能使用这个后缀来标识私钥文件。 `private.pri`。

**`.crt`**：通常用于表示 X.509 证书文件。这个后缀通常用于证书文件，格式可以是 PEM 或 DER。`server.crt`（服务器证书），`ca.crt`
（CA 证书）。

**`.csr`**：用于表示证书签名请求（Certificate Signing Request）文件。CSR 文件包含申请证书时的必要信息。`request.csr`。

**`.der`**：DER（Distinguished Encoding Rules）格式文件的标准后缀，通常用于存储二进制编码的证书或密钥。不同于 PEM 格式，DER
文件是二进制格式，没有 Base64 编码。`cert.der`（证书），`private.der`（私钥）。

## 操作

### （零）使用 OpenSSL 的流程

1. **生成密钥对**：使用 `openssl genpkey` 和 `openssl rsa` 生成私钥和公钥。
2. **创建 CSR**：使用 `openssl req` 创建证书签名请求。
3. **获得数字证书**：通过 CA 颁发或自签名生成证书。
4. **加密和签名**：使用公钥加密数据，使用私钥签名数据。
5. **解密和验证签名**：使用私钥解密数据，使用公钥验证签名。

### （一）安装OpenSSL

```bash
sudo apt install openssl
openssl version
```

### （二）生成自签名证书

生成一张自签名证书，类似一张独一无二的会员卡。

1. **生成私钥** 私钥是证书的核心。首先生成一个2048位的RSA私钥：

   ```bash
   # openssl genpkey: OpenSSL命令，用于生成私钥
   # algorithm RSA: 指定使用RSA算法生成私钥
   # out private.key: 输出文件名为private.key【可选，不写的话就是直接输出到终端】
   # pkeyopt rsa_keygen_bits:2048: 指定生成2048位的RSA密钥
   openssl genpkey -algorithm RSA -out private.key -pkeyopt rsa_keygen_bits:2048
   
   # 直接使用genrsa命令，就不用指定算法
   openssl genrsa -out genrsa.txt 512
   
   # 这个时候屏幕会出现一堆+.符号，这只是用来表示计算进度的输出
   # 查看私钥内容，是Base64编码格式
   cat private.key
   
   # 解析私钥，输出私钥的详细信息，包括私钥的模数（modulus）等
   openssl pkey -in private.key -text -noout
   
   # check检查私钥文件是否被修改过，如果修改了会输出RSA key not ok
   openssl rsa -in private.key -check
   
   # 输出私钥明文到文件夹中，就是把解析后的私钥输出
   openssl rsa -in private.key -text -out private.txt
   ```

   （可选）加钥私密

   ```bash
   # 加密私钥文件，加密的密码为123456
   openssl genrsa -out genrsa.txt -des3 -passout pass:123456 512
   
   # 密码错误，无法加密，必须一致才能加密
   openssl rsa -in genrsa.txt -passin pass:13456 -check
   ```


2. **生成证书签名请求（CSR）** 使用私钥生成一个CSR文件，这是获得数字证书的关键步骤，**CSR包含了申请者的身份信息和公钥**

   ```bash
   # openssl req: OpenSSL命令，用于处理证书请求和生成CSR。
   # new: 生成新的CSR。
   # key private.key: 使用之前生成的私钥private.key。
   # out mycsr.csr: 输出文件名为mycsr.csr。
   openssl req -new -key private.key -out mycsr.csr
   
   # 查看CSR内容，包含Subject（组织名称等）、Public Key（公钥）、Signature（用私钥对CSR进行的数字签名，确保CSR未被篡改）
   openssl req -in mycsr.csr -noout -text
   ```

   在这个过程中，会被提示输入一些信息，以下字段是必填的：

   1. Country Name (2 letter code):CN
   2. State or Province Name (full name):Guangdong
   3. Locality Name (eg, city):Guangzhou
   4. Organization Name (eg, company):SCNU
   5. Common Name (e.g. server FQDN or YOUR name):Xuan

   而以下字段是可选的：

   1. Organizational Unit Name (eg, section)
   2. Email Address
   3. A challenge password
   4. An optional company name

   ![image-20240725085508784](https://cdn.jsdelivr.net/gh/kixuan/PicGo/image-20240725085508784.png)

3. **生成自签名证书** 使用CSR和私钥生成一个有效期为一年的自签名证书：

   ```bash
   # penssl x509: OpenSSL命令，用于管理X.509证书。
   # req: 表示输入是一个证书请求（CSR）。
   # days 365: 证书的有效期为365天。
   # in mycsr.csr: 输入CSR文件mycsr.csr。
   # signkey private.key: 使用private.key私钥进行签名。
   openssl x509 -req -days 365 -in mycsr.csr -signkey private.key -out mycert.crt
   ```

### （三）基本使用

1. **查看证书信息** 使用以下命令查看证书的详细信息：

   ```bash
   # in mycert.crt: 输入证书文件mycert.crt。
   # text: 以文本格式显示证书内容。
   # noout: 不输出证书本身，只显示信息。
   # 数字证书包括版本号、序列号、签名算法、颁发者信息、有效期、持有者信息、公钥、证书扩展、签名等等
   openssl x509 -in mycert.crt -text -noout
   ```

2. **加密文件** 使用公钥加密一个文件：

   ```bash
   # 从证书中提取公钥
   openssl x509 -pubkey -noout -in mycert.crt > public.key
   
   # 生成测试文本文件，输入sth
   touch plaintext.txt
   / vim plaintext.txt
   
   # openssl rsautl: OpenSSL命令，用于使用RSA算法进行加密、解密操作。
   # encrypt: 进行加密操作。
   # inkey public.key: 使用public.key中的公钥。
   # pubin: 指定输入文件是公钥。
   # in plaintext.txt: 输入要加密的文件plaintext.txt。
   openssl pkeyutl -encrypt -pubin -inkey public.key -in plaintext.txt -out encrypted.dat
   ```

3. **解密文件** 使用私钥解密文件：

   ```bash
   # openssl pkeyutl: 用于使用RSA算法进行加密、解密操作。【注意3.0.x版本只能用pkeyutl不能用rsault】
   # decrypt: 进行解密操作。
   # inkey private.key: 使用私钥private.key进行解密。
   openssl pkeyutl -decrypt -inkey private.key -in encrypted.dat -out decrypted.txt
   
   # 查看解密后内容，发现和
   cat decrypted.txt
   ```

### （四）其他使用

1. `openssl speed`测试加密算法的性能

   ```bash
   openssl speed [md2] [mdc2] [md5] [hmac] [sha1] [rmd160] [idea-cbc] [rc2-cbc] [rc5-cbc] [bf-cbc] [des-cbc] [des-ede3] [rc4] [rsa512] [rsa1024] [rsa2048] [rsa4096] [dsa512] [dsa1024] [dsa2048] [idea] [rc2] [des] [rsa] [blowfish]
   ```

2. `openssl rand`生成伪随机数

   ```bash
   openssl rand [-out file] [-rand file(s)] [-base64] [-hex] num
   
   # 选项说明：
   # -out：指定随机数输出文件，否则输出到标准输出。
   # -rand file：指定随机数种子文件。种子文件中的字符越随机，openssl rand生成随机数的速度越快，随机度越高。
   # -base64：指定生成的随机数的编码格式为base64。
   # -hex：指定生成的随机数的编码格式为hex。
   # - num：指定随机数的长度。
   
   openssl rand -base64 30;
   openssl rand -hex 30;
   openssl rand 30
   ```

3. `openssl passwd` 生成加密的密码

   ```bash
   openssl passwd [-crypt] [-1] [-apr1] [-salt string] [-in file] [-stdin] [-quiet] {password}
   
   # 生成默认的密码哈希（使用 DES 加密）
   openssl passwd -1 "password"
   # 生成使用 SHA-256 的密码哈希
   openssl passwd -6 "password"
   # 生成盐（salt）并使用其生成密码哈希【盐相同，密码相同】
   openssl passwd -salt mysalt -1 "password"
   # 交互式输入密码
   openssl passwd -1

4. `openssl dgst` 生成和验证数字签名

   ```bash
   openssl dgst [-md5|-sha1|...] [-hex | -binary] [-out filename] [-sign filename] [-passin arg] [-verify filename] [-prverify filename] [-signature filename] [file...]
   
   # -sign filename:使用指定的私钥文件对哈希值进行签名
   # -verify filename:使用指定的公钥文件来验证签名
   # -prverify filename:同上，可以指定密钥文件的密码。
   # -signature filename:指定签名文件，以便与数据一起用于验证签名
   
   # 计算文件.txt 的 SHA-256 哈希并将结果写入 hash.txt 文件：
   openssl dgst -sha256 -hex (-out hash.txt) plaintext.txt 
   
   # 使用私钥对plaintext.txt生成签名，并将签名保存到signature.sig 中【使用-sign签名时绝对不能加-hex，不然后面验签一定会失败】
   openssl dgst -sha256 -sign private.key -out signature.sig plaintext.txt
   
   # 使用公钥 public.key 验证签名 signature.sig，输出verified.ok即签名有效
   openssl dgst -sha256 -verify public.key -signature signature.sig plaintext.txt 
   
   
   # 使用密码保护的私钥
   openssl dgst -sha256 -sign private.key -passin pass:yourpassword -out signature.sig example.txt
   
   ```

5. `openssl rsautl和openssl pkeyutl`：文件的非对称加密

   rsautl是rsa的工具，相当于rsa、dgst的部分功能集合，可用于生成数字签名、验证数字签名、加密和解密文件。

   pkeyutl是非对称加密的通用工具，大体上和rsautl的用法差不多

   ```bash
   ```

## 项目实践

### （一）在SpringBoot中使用

>
使用openssl创建自签名证书，并在SpringBoot应用中配置HTTPS，确保与HTTP请求兼容：[openssl创建自认证证书后，添加在springboot配置中，从而发起https请求，要求http和https请求兼容_csr证书怎么添加到springboot项目中去-CSDN博客](https://blog.csdn.net/csdnfzp2016/article/details/105154863)

1. OpenSSL生成密钥和证书，配置HTTPS

```yaml
# application.properties
server:
   ssl:
      key-store: classpath:keystore.p12
      key-store-password: yourpassword
      key-store-type: PKCS12
      key-alias: youralias
```

2.
在JWT中使用（JWT详细介绍可见：[浅析JWT原理及牛客出现过的相关面试题-CSDN博客](https://blog.csdn.net/Kixuan214/article/details/140690612?spm=1001.2014.3001.5502)
），在Spring Boot应用中使用JWT进行认证和授权时，可以使用生成的私钥和公钥进行签名和验证。

```java
import io.jsonwebtoken.Jwts;
import io.jsonwebtoken.SignatureAlgorithm;
import java.security.KeyFactory;
import java.security.PrivateKey;
import java.security.PublicKey;
import java.security.spec.PKCS8EncodedKeySpec;
import java.security.spec.X509EncodedKeySpec;
import java.util.Base64;

public class JwtUtil {

    private static final String PRIVATE_KEY = "-----BEGIN PRIVATE KEY-----...-----END PRIVATE KEY-----";
    private static final String PUBLIC_KEY = "-----BEGIN PUBLIC KEY-----...-----END PUBLIC KEY-----";

    // 从PEM格式的字符串中解析并返回私钥
    private static PrivateKey getPrivateKey() throws Exception {
        String privateKeyPEM = PRIVATE_KEY.replace("-----BEGIN PRIVATE KEY-----", "")
                                          .replace("-----END PRIVATE KEY-----", "")
                                          .replaceAll("\\s", "");
        byte[] encoded = Base64.getDecoder().decode(privateKeyPEM);
        PKCS8EncodedKeySpec keySpec = new PKCS8EncodedKeySpec(encoded);
        KeyFactory keyFactory = KeyFactory.getInstance("RSA");
        return keyFactory.generatePrivate(keySpec);
    }

    // 从PEM格式的字符串中解析并返回公钥
    private static PublicKey getPublicKey() throws Exception {
        String publicKeyPEM = PUBLIC_KEY.replace("-----BEGIN PUBLIC KEY-----", "")
                                        .replace("-----END PUBLIC KEY-----", "")
                                        .replaceAll("\\s", "");
        byte[] encoded = Base64.getDecoder().decode(publicKeyPEM);
        X509EncodedKeySpec keySpec = new X509EncodedKeySpec(encoded);
        KeyFactory keyFactory = KeyFactory.getInstance("RSA");
        return keyFactory.generatePublic(keySpec);
    }

    // 使用私钥生成JWT令牌
    public static String generateToken(String subject) throws Exception {
        return Jwts.builder()
                   .setSubject(subject)
                   .signWith(getPrivateKey(), SignatureAlgorithm.RS256)
                   .compact();
    }

    // 使用公钥验证JWT令牌
    public static boolean validateToken(String token) throws Exception {
        Jwts.parserBuilder()
            .setSigningKey(getPublicKey())
            .build()
            .parseClaimsJws(token);
        return true;
    }
}
```

### （二）在区块链项目中使用

使用私钥对交易进行签名。使用Java的Bouncy Castle库进行签名

1. 添加Bouncy Castle依赖

   ```pom
   <dependency>
       <groupId>org.bouncycastle</groupId>
       <artifactId>bcprov-jdk15on</artifactId>
       <version>1.70</version>
   </dependency>
   <dependency>
       <groupId>org.bouncycastle</groupId>
       <artifactId>bcpkix-jdk15on</artifactId>
       <version>1.70</version>
   </dependency>
   ```

2. 创建证书生成工具类`CertificateGenerator`

   ```java
   import org.bouncycastle.asn1.x500.X500Name;
   import org.bouncycastle.asn1.x500.style.BCStyle;
   import org.bouncycastle.asn1.x509.Extension;
   import org.bouncycastle.asn1.x509.ExtensionsGenerator;
   import org.bouncycastle.cert.X509CertificateHolder;
   import org.bouncycastle.cert.jcajce.JcaX509CertificateConverter;
   import org.bouncycastle.cert.jcajce.JcaX509v3CertificateBuilder;
   import org.bouncycastle.jce.provider.BouncyCastleProvider;
   import org.bouncycastle.operator.ContentSigner;
   import org.bouncycastle.operator.jcajce.JcaContentSignerBuilder;
   import org.bouncycastle.operator.jcajce.JcaDigestCalculatorProviderBuilder;
   
   import java.math.BigInteger;
   import java.security.*;
   import java.security.cert.CertificateException;
   import java.security.cert.X509Certificate;
   import java.util.Date;
   
   public class CertificateGenerator {
   
       static {
           Security.addProvider(new BouncyCastleProvider());
       }
   
       public static KeyPair generateKeyPair() throws NoSuchAlgorithmException {
           KeyPairGenerator keyPairGenerator = KeyPairGenerator.getInstance("RSA");
           keyPairGenerator.initialize(2048);
           return keyPairGenerator.generateKeyPair();
       }
   
       public static X509Certificate generateCertificate(KeyPair keyPair) throws Exception {
           X500Name issuer = new X500Name("CN=Test CA Certificate");
           X500Name subject = new X500Name("CN=Test Certificate");
           BigInteger serialNumber = BigInteger.valueOf(System.currentTimeMillis());
           Date notBefore = new Date(System.currentTimeMillis() - 24 * 60 * 60 * 1000);
           Date notAfter = new Date(System.currentTimeMillis() + 365 * 24 * 60 * 60 * 1000L);
           PublicKey publicKey = keyPair.getPublic();
           
           JcaX509v3CertificateBuilder builder = new JcaX509v3CertificateBuilder(
                   issuer, serialNumber, notBefore, notAfter, subject, publicKey);
   
           ExtensionsGenerator extensionsGenerator = new ExtensionsGenerator();
           extensionsGenerator.addExtension(Extension.basicConstraints, true, new org.bouncycastle.asn1.x509.BasicConstraints(true));
           builder.addExtension(Extension.basicConstraints, true, new org.bouncycastle.asn1.x509.BasicConstraints(true));
   
           ContentSigner contentSigner = new JcaContentSignerBuilder("SHA256withRSA").build(keyPair.getPrivate());
   
           X509CertificateHolder certificateHolder = builder.build(contentSigner);
           return new JcaX509CertificateConverter().setProvider("BC").getCertificate(certificateHolder);
       }
   }
   ```

3. 使用工具类`CertificateGenerator`生成密钥对和自签名证书

   ```java
   import org.springframework.boot.CommandLineRunner;
   import org.springframework.boot.SpringApplication;
   import org.springframework.boot.autoconfigure.SpringBootApplication;
   
   import java.io.FileOutputStream;
   import java.security.KeyPair;
   import java.security.cert.X509Certificate;
   
   @SpringBootApplication
   public class SpringBootCertificateApp implements CommandLineRunner {
   
       public static void main(String[] args) {
           SpringApplication.run(SpringBootCertificateApp.class, args);
       }
   
       @Override
       public void run(String... args) throws Exception {
           KeyPair keyPair = CertificateGenerator.generateKeyPair();
           X509Certificate certificate = CertificateGenerator.generateCertificate(keyPair);
   
           try (FileOutputStream keyOut = new FileOutputStream("private.key")) {
               keyOut.write(keyPair.getPrivate().getEncoded());
           }
   
           try (FileOutputStream certOut = new FileOutputStream("certificate.crt")) {
               certOut.write(certificate.getEncoded());
           }
   
           System.out.println("Certificate and private key generated successfully.");
       }
   }
   
   ```

## TIPS/BUGS

### 验证openssl version时报错version not found

背景：在使用openssl
version时出现`openssl: /lib/x86_64-linux-gnu/libcrypto.so.3: version OPENSSL_3.0.9' not found (required by openssl)`

解决办法：

```bash
# 1. 更新
sudo apt upgrade openssl

# 2. 重装
sudo apt remove --purge openssl
sudo apt install openssl
```

### 加密时报错`Could not read public key from mycert.crt`

意思是在证书中读取不到公钥，这种情况就把公钥提取出来，再用公钥文件加密即可

``` bash
# 从证书中提取公钥
openssl x509 -pubkey -noout -in mycert.crt > public.key

# 加密
openssl pkeyutl -encrypt -pubin -inkey public.key -in plaintext.txt -out encrypted.dat
```

### `openssl genrsa` 和`openssl genpkey -algorithm rsa`有什么区别?

openssl pkey 更通用，可以处理多种私钥
openssl rsa 专门处理RSA密钥

[openssl-genpkey - OpenSSL Documentation](https://docs.openssl.org/master/man1/openssl-genpkey/#notes)
明确指出应使用`genpkey`，而不是特定于算法的 `genrsa`：

> The use of the genpkey program is encouraged over the algorithm specific utilities because additional algorithm
> options and ENGINE provided algorithms can be used.

## 实现

> 基于区块链平台Hyperledger
> Fabric，参考官方示例实现一个智能合约，同时基于SpringBoot和官方SDK实现一个应用程序，用于实现数字证书的存储、查询等；结合上述内容，使用OpenSSL为现有系统实现数字证书的生成、私钥签名、验签等内容。

### **Hyperledger Fabric 智能合约实现**

该智能合约用于存储和管理数字证书，需要实现的功能包括：

- 创建新证书：将证书信息写入区块链。
- 查询证书：从区块链中读取证书信息。
- 更新证书：修改证书信息。
- 撤销证书：标记证书无效。

### **OpenSSL 数字证书管理**

使用 OpenSSL 为现有系统实现数字证书的生成、私钥签名、验签等功能,该部分与 Hyperledger Fabric 集成，用于：

- 生成证书：为新用户生成数字证书。
- 签名：使用私钥对证书进行签名。
- 验签：验证证书的有效性。
