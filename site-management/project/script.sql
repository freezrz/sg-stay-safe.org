--
-- Create model Rule
--
CREATE TABLE `adminportal_rule` (`id` bigint AUTO_INCREMENT NOT NULL PRIMARY KEY, `name` varchar(200) NOT NULL, `description` varchar(200) NOT NULL, `is_enabled` bool NOT NULL, `value` integer NOT NULL);
--
-- Create model Site
--
CREATE TABLE `adminportal_site` (`id` bigint AUTO_INCREMENT NOT NULL PRIMARY KEY, `name` varchar(200) NOT NULL, `site_id` varchar(50) NOT NULL, `address` varchar(200) NOT NULL, `postal_code` varchar(10) NOT NULL, `description` varchar(200) NOT NULL, `capacity` integer NOT NULL, `should_ban` bool NOT NULL);
