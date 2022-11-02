<?php
/**
 * Create By: Will Yin
 * Date: 2021/2/28
 * Time: 16:13
 **/
namespace Design\Decorator;

abstract class FoodDecorator implements Food {

    /**
     * @var $decoratorFood
     */
    protected $decoratorFood;

    public function __construct(Food $food) {
        $this->decoratorFood = $food;
    }

    public function make() {
        $this->decoratorFood->make();
    }

}