<?php

declare(strict_types=1);

namespace services;

final class View
{
    private const XML_TYPE = 'xml';
    private const HTML_TYPE = 'html';
    //private const JSON_TYPE = 'json';

    private const ALLOW_TYPES = [self::XML_TYPE, self::HTML_TYPE/*, self::JSON_TYPE*/];

    private string $type;

    public function __construct(string $type = self::XML_TYPE)
    {
        if (! in_array($type, self::ALLOW_TYPES, true)) {
            throw new \RuntimeException("Неизвестный тип для вывода: $type");
        }

        $this->type = $type;
    }

    public function __call(string $view, array $args)
    {
        echo strtr(require __DIR__ . "/../views/{$this->type}/{$view}.php", $args[0] ?? []);
    }
}