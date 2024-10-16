<?php

declare(strict_types=1);

namespace services;

final class View
{
    public function __call(string $view, array $args)
    {
        echo strtr(require __DIR__ . "/../views/xml/{$view}.php", $args[0] ?? []);
    }
}