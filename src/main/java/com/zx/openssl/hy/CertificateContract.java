package com.zx.openssl.hy;

import io.swagger.annotations.Info;
import net.bytebuddy.implementation.bind.annotation.Default;
import org.apache.http.annotation.Contract;
import org.hyperledger.fabric.protos.peer.ChaincodeGrpc;
import org.hyperledger.fabric.protos.peer.ChaincodeID;
import org.hyperledger.fabric.protos.peer.ProposalResponse;

import java.nio.charset.StandardCharsets;
import java.util.Collection;

@Contract(
        name = "CertificateContract",
        info = @Info(
                title = "Certificate contract",
                description = "A simple chaincode for managing certificates",
                version = "0.0.1"
        ))
@Default
public class CertificateContract implements ContractInterface {

    @Transaction()
    public void createCertificate(Context ctx, String id, String certificate) {
        ChaincodeGrpc.ChaincodeStub stub = ctx.getStub();

        if (stub.getState(id).length != 0) {
            String errorMessage = String.format("Certificate with id %s already exists", id);
            throw new ChaincodeException(errorMessage, errorMessage);
        }

        stub.putStringState(id, certificate);
    }

    @Transaction()
    public String getCertificate(Context ctx, String id) {
        ChaincodeStub stub = ctx.getStub();
        byte[] certificateBytes = stub.getState(id);

        if (certificateBytes.length == 0) {
            String errorMessage = String.format("Certificate with id %s does not exist", id);
            throw new ChaincodeException(errorMessage, errorMessage);
        }

        return new String(certificateBytes, StandardCharsets.UTF_8);
    }

    // 调用智能合约
    HFClient client = HFClient.createNewInstance();
// ... 初始化 client 并连接到 Fabric 网络 ...

    Channel channel = client.getChannel("mychannel");
    ChaincodeID chaincodeID = ChaincodeID.newBuilder().setName("certificate_cc").build();

    // 创建证书
    CreateCertificateProposalResponse createResponse = channel.sendTransactionProposal(
            client.newTransactionProposalRequest()
                    .setChaincodeID(chaincodeID)
                    .setFcn("createCertificate")
                    .setArgs(id, certificate)
    );

    // 查询证书
    QueryByChaincodeRequest queryRequest = client.newQueryProposalRequest()
            .setChaincodeID(chaincodeID)
            .setFcn("getCertificate")
            .setArgs(id);
    Collection<ProposalResponse> queryResponses = channel.queryByChaincode(queryRequest);

    // ... 处理响应 ...
    @Transaction()
    public void deleteCertificate(Context ctx, String id) {
        ChaincodeStub stub = ctx.getStub();
        stub.delState(id);
    }
}