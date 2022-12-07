CREATE TABLE `tickers` (
    `symbol` text NOT NULL,
    `created_at` datetime,
    `updated_at` datetime
    ,PRIMARY KEY (`symbol`)
);

CREATE TABLE `aggregates` (
    `id` integer,
    `symbol` text NOT NULL,
    `price` real,
    `date` datetime,
    PRIMARY KEY (`id`),
    CONSTRAINT `fk_tickers_prices` FOREIGN KEY (`symbol`) REFERENCES `tickers`(`symbol`)
);

CREATE INDEX `symbol_date` ON `aggregates`(`symbol`,`date` desc);
