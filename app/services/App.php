<?php

declare(strict_types=1);

namespace services;

final class App
{
    private const ALLOW_EXTENSIONS = [
        'avi',
        'mkv',
        'mp4',
        'mpeg',
        'mpg',
        'ts',
        'vob',
        'wmv',
        'wvc',
        'webm',
        'asf',
        '3gp',
        '3g2',
        'mts',
        'mov',
        'flv',
    ];

    private View $view;

    private string $source;
    private string $fullSource;
    private string $fullHost;
    private bool $isTv;

    public function __invoke(): void
    {
        error_reporting(E_ALL);

        $this->params();

        $this->delete();
        $this->header($this->isIncorrectInput());
        $this->view();
    }

    private function params(): void
    {
        $this->isTv = $this->isTv();
        $this->source = $this->getSource();
        $this->fullSource = $_SERVER['DOCUMENT_ROOT'] . '/data/' . (empty($this->source) ? $this->source : $this->source . '/');
        $this->fullHost = $_SERVER['REQUEST_SCHEME'] . '://' . $_SERVER['HTTP_HOST'];

        require_once 'View.php';
        $this->view = new View($this->isTv ? 'xml' : 'html');
    }

    private function header(bool $hasRedirect = false, string $location = null): void
    {
        if ($hasRedirect) {
            header('HTTP/1.1 301 Moved Permanently');
            header("Location: " . ($location ?? $this->fullHost));
            exit();
        }

        header($this->isTv ? "Content-Type: text/xml" : "Content-Type: text/html");
        header("Expires: Thu, 19 Feb 1998 13:24:18 GMT");
        header("Last-Modified: " . gmdate("D, d M Y H:i:s") . " GMT");
        header("Cache-Control: no-cache, must-revalidate");
        header("Cache-Control: post-check=0,pre-check=0");
        header("Cache-Control: max-age=0");
        header("Pragma: no-cache");
    }

    private function view(): void
    {
        $this->showTop();
        $this->showPreviousDirs();
        $this->showItems(...(isset($_GET['search']) ? $this->getSearch() : $this->getList()));
        $this->view->bottom();
    }

    ### Визуализация ###

    private function showTop(): void
    {
        $this->view->top(['{title}' => empty($this->source) ? 'Корень' : $this->source, '{fullHost}' => $this->fullHost]);
    }

    private function showPreviousDirs(): void
    {
        if ('' === $this->source) {
            return;
        }

        $this->view->root_dir(['{fullHost}' => $this->fullHost]);

        if (! empty(pathinfo($this->source)['dirname']) && '.' !== pathinfo($this->source)['dirname']) {
            $this->view->previous_dir(['{prev}' => $this->fullHost . '?source=' . pathinfo($this->source)['dirname']]);
        }
    }

    private function showItems(array $dirs, array $files): void
    {
        foreach ($dirs as $dir) {
            $fullPath = empty($this->source) ? $dir : $this->source . '/' . $dir;
            $this->view->playlist(['{dir}' => $dir, '{fullHost}' => $this->fullHost, '{fullPath}' => $fullPath]);
        }

        foreach ($files as $key => $file) {
            $fullPath = is_int($key) ? (empty($this->source) ? $file : $this->source . '/' . $file) : $key;
            $this->view->stream(['{file}' => $file, '{fullHost}' => $this->fullHost, '{fullPath}' => $fullPath, '{source}' => $this->source]);
        }
    }

    ### Вспомогательное ###

    private function isIncorrectInput(): bool
    {
        return ! empty($_GET['source']) && $this->source !== $_GET['source'];
    }

    private function getSource(): string
    {
        $source = ($_GET['source'] ?? null);

        if (null === $source || '/' === $source || '..' === $source || '.' === $source) {
            return '';
        }

        if ('/' === $source[0]) {
            $source = substr($source, 1);
        }

        $source = str_replace(['../', './'], '', $source);

        return empty($source) ? '' : $source;
    }

    private function getList(): array
    {
        $list = scandir($this->fullSource);
        $dirs = [];
        $files = [];

        foreach ($list as $item) {
            if (in_array($item, ['..', '.'], true)) {
                continue;
            }

            if (is_dir($this->fullSource . $item)) {
                $dirs[] = $item;
                continue;
            }

            if (in_array(mb_strtolower(pathinfo($item, PATHINFO_EXTENSION)), self::ALLOW_EXTENSIONS, true)) {
                $files[] = $item;
            }
        }

        return [$dirs, $files];
    }

    private function getSearch(): array
    {
        $files = [];

        $scan = function ($source) use (&$scan, &$files) {
            $list = scandir($source);
            foreach ($list as $item) {
                if (in_array($item, ['..', '.'], true)) {
                    continue;
                }

                $currentSource = $source . '/' . $item;

                if (is_dir($currentSource)) {
                    $scan($currentSource);
                    continue;
                }

                if (str_contains(mb_strtolower($item), mb_strtolower($_GET['search']))) {
                    $files[substr($currentSource, mb_strlen($this->fullSource))] = $item;
                }
            }
        };

        $scan(substr($this->fullSource, 0, -1));

        return [[], $files];
    }

    private function isTv(): bool
    {
        $words = ['TV', 'Tizen', 'webOS', 'Chromecast', 'Fork'];
        foreach ($words as $word) {
            if (str_contains($_SERVER['HTTP_USER_AGENT'], $word)) {
                return true;
            }
        }

        return false;
    }

    private function delete(): void
    {
        if (isset($_GET['delete'])) {
            @unlink($this->fullSource . $_GET['delete']); // LASTEDIT Нужно удалять и директории
            $this->header(true, "/?source={$this->source}");
        }
    }
}
