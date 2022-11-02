<?php
/**
 * Create By: Will Yin
 * Date: 2021/2/28
 * Time: 12:20
 **/
namespace Design\Simple\Register;

class Index {

    /**
     * @var array $objects
     */
    protected static $objects;

    public static function set($alias,$object) {
        self::$objects[$alias] = $object;
    }

    public static function __unset($alias) {
        unset(self::$objects[$alias]);
    }

    public static function get($alias) {
        return self::$objects[$alias];
    }
}