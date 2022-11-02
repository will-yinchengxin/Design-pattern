<?php
/**
 * Create By: Will Yin
 * Date: 2021/2/28
 * Time: 15:18
 **/
namespace Design\Prototype;

class Book {
    /**
     * @var $time
     */
    private $time;

    /**
     * @var $name
     */
    private $name;

    /**
     * @var $price
     */
    private $price;

    public function setTime()
    {
        $this->time = date("Y-m-d");
    }

    public function setName(string $name)
    {
        $this->name = $name;
    }

    public function setPrice()
    {
        $this->price = 25;
    }

    public function getInfo()
    {
        echo $this->name."在".$this->time."在本图书馆借阅了图书,价格为".$this->price.PHP_EOL;
    }

}