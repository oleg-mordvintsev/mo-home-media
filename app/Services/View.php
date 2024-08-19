<?php

declare(strict_types=1);

namespace Services;

final class View
{
    public function __call(string $view, array $args)
    {
        echo strtr(require __DIR__ . "/../Views/{$view}.php", $args[0] ?? []);
    }
}