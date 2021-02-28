<?php
/**
 * Create By: Will Yin
 * Date: 2021/2/28
 * Time: 11:36
 **/

// 引入自动加载文件
require __DIR__.'/vendor/autoload.php';

// 单例设计模式
//use Design\Simple\Single\Index as Single;
//Single::getInstance();


// 工厂模式
//use Design\Simple\Factory\Index as Factory;
//echo Factory::getData();


// 注册模式
//use Design\Simple\Register\Index as Register;
// 可以在项目执行的任意时刻进行注册
//Register::get('data');


// 适配器模式(同一规则可以无缝切换)
//$connect = (new \Design\Adapter\Database\Mysql())->connect('localhost', 6379);
//echo $connect->query();
//$connect = (new \Design\Adapter\Database\Pdo())->connect('localhost', 6380);
//echo $connect->query();


// 策略模式
//use Design\Strategy\Index as Strategy;
//($page = new Strategy())->judgeStrategy($_GET);
//$page->page();


// 观察者模式
//use Design\Observer\GenderEvent;
//// 添加观察者
//($event = new GenderEvent())->addObserver(new \Design\Observer\SchoolObserver());
//$event->addObserver(new \Design\Observer\ShopObserver());
//$event->trigger();


// 原型模式(克隆带来的问题: https://blog.csdn.net/qq_43314956/article/details/105210515)
//$book = new \Design\Prototype\Book();
//// 小明信息
//$book->setName('小明');
//$book->getInfo();
//// 小红信息
//$book_one = clone $book;
//$book_one->setName('小红');
//$book_one->getInfo();


// 装饰器模式
//$food = new \Design\Decorator\Cake();
//$foodDecorator = new \Design\Decorator\NoodleFoodDecorator($food);
//$foodDecorator->make();


// 迭代器模式(项目中极少使用)
//$info = new \Design\Iterator\Index();
//foreach ($info as $key => $value) {
//    $info->add_time = time();
//    $info->save();
//}


// 代理模式
//$proxy = new \Design\Proxy\UserProxy();
//$proxy->getName(1);
//$proxy->updateName(1,'will');