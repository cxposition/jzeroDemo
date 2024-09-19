CREATE TABLE `log`
(
    `uuid` VARCHAR(191) NOT NULL,
    `operator` VARCHAR(191) NOT NULL,
    `ipaddr` VARCHAR(191) NOT NULL,
    `actions` VARCHAR(191) NOT NULL,
    `contents` VARCHAR(191) NOT NULL,
    `status`   bigint NOT NULL,
    `create_time` bigint NOT NULL,
    `event_type`bigint NOT NULL,
    `event_level` bigint NOT NULL,
    `audit_status` bigint NOT NULL,
    PRIMARY KEY(`uuid`)
)
ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci;