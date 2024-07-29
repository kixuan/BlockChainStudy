package com.zx.openssl;

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
