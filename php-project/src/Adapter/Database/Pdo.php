<?php
/**
 * Create By: Will Yin
 * Date: 2021/2/28
 * Time: 13:53
 **/
namespace Design\Adapter\Database;

class Pdo implements Datebase {

    function connect($host, $port)
    {
        echo "Host 为".$host.",Port为".$port;
        return $this;
    }

    function query()
    {
        echo "这里进行pdo连接";
    }

}