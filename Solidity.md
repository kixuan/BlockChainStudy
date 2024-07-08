一个简单的合约例子，用于存储键值对，并提供基本的增删查改操作

```solidity
pragma solidity ^0.5.0;  // 指定Solidity版本
pragma experimental ABIEncoderV2;  //允许使用实验性的ABI编码

import "./AbstractBean.sol";
import "./LibStrings.sol";

/**
 * @title NFT Map 合约
 * 定义 MapStorage 合约，继承 AbstractBean【基础合约类】
 */
contract MapStorage is AbstractBean {
    string privKey;
    string uniKey;
    string[] mFields;

    // 构造函数
    // 第一个参数是表名称，用于识别和区分不同表
    constructor() public AbstractBean("mapStorage", "id", "key", "value") {
        privKey = "id";
        uniKey = "key";
        mFields = splitString("value");
    }

    // 创建键值对
    // require 条件检查和错误处理
    // memory 关键字表示变量只存在于函数调用期间，而不会被永久存储在区块链上，临时的，函数执行完后会被销毁
    function createMap(string memory key, string memory value) public {
        (int res, string memory result) = select(key);
        require(res != 0, "key exists!");
        string[] memory fields = new string[](1);
        fields[0] = value;
        insert(key, key, fields);
    }

    // 读取键值对
    function readMap(string memory key) public view returns (int, string memory) {
        (int res, string memory result) = select(key);
        return (res, result);
    }

    // 检查键是否存在
    function existMap(string memory key) public view returns (int) {
        int res = exist(key);
        return res;
    }

    // 更新键值对
    function updateMap(string memory key, string memory value) public {
        (int res, string memory result) = select(key);
        require(res == 0, "key does not exist!");
        string[] memory values = new string[](1);
        values[0] = value;
        update(key, mFields, values);
    }

    // 删除键值对
    function deleteMap(string memory key) public {
        (int res,) = select(key);
        require(res == 0, "key not found");
        remove(key);
    }

    // 分割字符串   把value字符串拆分成包含value的数组
    // internal 限制访问权限，只能在当前和继承的合约调用
    // pure 表示该函数不读取也不修改区块链上的状态（即合约的存储状态）
    function splitString(string memory _value) internal pure returns (string[] memory) {
        LibStrings.slice memory s = _value.toSlice();
        LibStrings.slice memory delim = ",".toSlice();
        string[] memory values = new string[](s.count(delim) + 1);
        for (uint i = 0; i < values.length; i++) {
            values[i] = s.split(delim).toString();
        }
        return values;
    }
}

```

