package com.zx.openssl;

import org.bouncycastle.asn1.x500.X500Name;
import org.bouncycastle.asn1.x509.Extension;
import org.bouncycastle.asn1.x509.ExtensionsGenerator;
import org.bouncycastle.cert.X509CertificateHolder;
import org.bouncycastle.cert.jcajce.JcaX509CertificateConverter;
import org.bouncycastle.cert.jcajce.JcaX509v3CertificateBuilder;
import org.bouncycastle.jce.provider.BouncyCastleProvider;
import org.bouncycastle.operator.ContentSigner;
import org.bouncycastle.operator.jcajce.JcaContentSignerBuilder;

import java.math.BigInteger;
import java.security.*;
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