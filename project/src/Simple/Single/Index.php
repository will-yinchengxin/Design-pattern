<?php
/**
 * Create By: Will Yin
 * Date: 2021/2/28
 * Time: 11:37
 **/
namespace Design\Simple\Single;

class Index {

    /**
     * @var $static
     */
    protected static $static;

    /**
     * 声明构造方法为私有
     */
    private function __construct() {

    }

    /**
     * 声明一个静态方法获取实例
     */
    public static function getInstance() {
        if (!self::$static) {
            self::$static = new self();
        }
        return self::$static;
    }

    // 封锁clone
    final private function __clone(){}

}