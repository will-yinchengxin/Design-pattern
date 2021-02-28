<?php
/**
 * Create By: Will Yin
 * Date: 2021/2/28
 * Time: 16:17
 **/
namespace Design\Decorator;

class NoodleFoodDecorator extends FoodDecorator {

    public function __construct(Food $food) {
        parent::__construct($food);
    }


    public function make() {
        $this->decoratorFood->make();
        $this->setNoodleFood($this->decoratorFood);
    }

    private function setNoodleFood(Food $food)
    {
        echo 'noodle';
    }

}