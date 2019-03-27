-- phpMyAdmin SQL Dump
-- version 4.5.4.1deb2ubuntu2
-- http://www.phpmyadmin.net
--
-- Host: localhost
-- Generation Time: Mar 26, 2019 at 07:08 AM
-- Server version: 5.7.18-0ubuntu0.16.04.1
-- PHP Version: 7.0.28-1+ubuntu16.04.1+deb.sury.org+1

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `taxexpress`
--

-- --------------------------------------------------------

--
-- Table structure for table `admins`
--

CREATE TABLE `admins` (
  `id` int(10) UNSIGNED NOT NULL,
  `username` varchar(64) DEFAULT NULL,
  `email` varchar(45) DEFAULT NULL,
  `token` varchar(255) DEFAULT NULL,
  `password` varchar(255) NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Dumping data for table `admins`
--

INSERT INTO `admins` (`id`, `username`, `email`, `token`, `password`, `created_at`, `updated_at`) VALUES
(1, 'admin', 'oshoklinsmann@gmail.com', '$2a$10$gfhsL7CYzSKzOBeWMpR4T.TsatYSzpWfU5ttoqU16hkfL/TbX7RhC', '$2a$04$x7KXAaIaxbUT.5SmgBKor.va8GcwGptfF16c4v4vDTr2qRxBP.rkm', '2019-02-27 22:29:26', '2019-02-27 22:29:26'),
(2, 'test_user', 'test@gmail.com', '$2a$10$nzmvD3cWJYWGhBVIBfiMiuVulOFVxdMjfvfDConm/oC/Y1n7ctVoO', '$2a$04$RoHy65saKb.YVLROT2HCQe8tOlfoE4qY.8IIARc0AgSx32fQc7dHq', '2019-03-16 13:48:16', '2019-03-16 13:48:16'),
(3, 'toluolu', 'toluolu72@gmail.com', '$2a$10$OAScdfq7R7L7c6zwgvbaceubKNxiGZf7GWxmpZ0fH7peVDhg7re/C', '$2a$04$./wRuooi1BzSIERkUooA.uTzad05uQDyBqeBXVsWI6CTpwDI28kri', '2019-03-16 14:48:56', '2019-03-16 14:48:56'),
(4, 'test_user2', 'testuser2@gmail.com', '$2a$10$ce1GQGdsELJdobhhsHClHOQFwn1tQrVA8BqT6wvrBYCdhNH0eNM5K', '$2a$04$vZG83OvVm1VnXPPvE1AraOhL9.pv.htkezmC5Z9lfCISIA3ZEkPs6', '2019-03-16 14:56:01', '2019-03-16 14:56:01'),
(5, 'test_user3', 'testuser3@gmail.com', '$2a$10$2b99rjxrAZE0xuhv2v/.9eqNbdg7zVGdhxzuLSGNgocnDTaPuDtNW', '$2a$04$Qg0XM61YBn7Dl26CmZ7YiegFaph8x.FxyIikdur.DNEXrT2ynK8jy', '2019-03-16 14:58:42', '2019-03-16 14:58:42'),
(6, 'test_user4', 'testuser4@gmail.com', '$2a$10$SkI46VxAkxkf0RRhFqyM3e1v5XceUoZGGvhsL/8QD3Z/xMIiFHk0S', '$2a$04$wfkMf5WtidzvgN07k1DQDuzPXyj6U6aF14juVvMXHvfPikssWflHO', '2019-03-16 15:19:06', '2019-03-16 15:19:06');

-- --------------------------------------------------------

--
-- Table structure for table `businesses`
--

CREATE TABLE `businesses` (
  `id` int(11) UNSIGNED NOT NULL,
  `name` varchar(45) DEFAULT NULL,
  `rc_number` int(8) DEFAULT NULL,
  `business_desc` varchar(255) DEFAULT NULL,
  `email` varchar(45) NOT NULL,
  `total_revenue` float DEFAULT NULL,
  `tax_status` int(11) NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Dumping data for table `businesses`
--

INSERT INTO `businesses` (`id`, `name`, `rc_number`, `business_desc`, `email`, `total_revenue`, `tax_status`, `created_at`, `updated_at`) VALUES
(1, 'Exampulse Educational Services', 234524, 'dealing in educational resources', 'info@exampulse.com', 45000, 1, '2019-02-26 15:15:25', '2019-02-26 16:17:38'),
(2, 'Prestavo', 234521, 'selling travel accessories', 'info@prestavo.com', 2500000, 1, '2019-02-26 15:24:38', '2019-02-26 15:25:39'),
(3, 'TheSocialmart', 204521, 'selling social media services', 'info@socialmart.com', 45000, 1, '2019-02-28 13:52:56', '2019-02-28 17:18:50'),
(4, 'Osho Enterprise', 234572, 'A business owned by Osho Emmanuel', 'oshoklinsmann@gmail.com', 0, 1, '2019-03-24 21:43:44', '2019-03-24 21:43:44'),
(6, 'Osho Enterprise', 234572, 'A business owned by Osho Emmanuel', 'kaywee@gmail.com', 0, 1, '2019-03-24 21:44:06', '2019-03-24 21:44:06'),
(7, 'Indoor Decor', 234578, 'A business about indoor decorations', 'indoordecor@gmail.com', 0, 1, '2019-03-24 22:11:46', '2019-03-24 22:11:46'),
(8, 'baby Store', 234579, 'Online Baby Store in Nigeria.', 'babystore@gmail.com', 0, 1, '2019-03-25 05:42:20', '2019-03-25 05:42:20'),
(9, 'Hotsauce', 234521, 'Eat the hottest pepper sauce chicken in Nigeria.', 'hotsauce@gmail.com', 0, 1, '2019-03-25 05:44:28', '2019-03-25 05:44:28'),
(10, 'Naijaloaded', 20234, 'Nigeria\'s number 1 news website', 'hello@naijaloaded.com', 0, 1, '2019-03-25 05:52:37', '2019-03-25 05:52:37'),
(11, 'Beatz Unplugged', 202224, 'Selling resources for music and sound production', 'hello@beatzunplugged.com', 0, 1, '2019-03-25 05:54:14', '2019-03-25 05:54:14'),
(12, 'Anakle ', 23487, 'A  digital agency building experiences for online and offline businesses.', 'anakle@gmail.com', 65000, 1, '2019-03-25 06:26:57', '2019-03-25 07:13:13');

-- --------------------------------------------------------

--
-- Table structure for table `messages`
--

CREATE TABLE `messages` (
  `id` int(10) UNSIGNED NOT NULL,
  `message` varchar(516) DEFAULT NULL,
  `type` varchar(45) DEFAULT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Dumping data for table `messages`
--

INSERT INTO `messages` (`id`, `message`, `type`, `created_at`, `updated_at`) VALUES
(1, 'you have pending taxes to be paid on your business', 'warning', '2019-03-17 00:44:52', '2019-03-17 00:44:52'),
(2, 'your unpaid taxes have been reported to appropriate authorities. It is advised that you see to this immediately.', 'critical', '2019-03-17 00:47:34', '2019-03-17 00:47:34');

-- --------------------------------------------------------

--
-- Table structure for table `notifications`
--

CREATE TABLE `notifications` (
  `id` int(10) UNSIGNED NOT NULL,
  `business_id` int(11) UNSIGNED NOT NULL,
  `message` varchar(255) NOT NULL,
  `to_business` tinyint(1) NOT NULL,
  `to_authority` tinyint(1) NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- --------------------------------------------------------

--
-- Table structure for table `tax`
--

CREATE TABLE `tax` (
  `id` int(10) UNSIGNED NOT NULL,
  `business_id` int(11) UNSIGNED NOT NULL,
  `tax_period` date NOT NULL,
  `revenue` float NOT NULL,
  `tax_paid` float NOT NULL DEFAULT '0',
  `date_paid` datetime DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Dumping data for table `tax`
--

INSERT INTO `tax` (`id`, `business_id`, `tax_period`, `revenue`, `tax_paid`, `date_paid`, `created_at`, `updated_at`) VALUES
(1, 1, '2019-02-25', 45000, 4500, '2019-04-25 00:00:00', '2019-02-26 15:15:39', '2019-02-26 16:19:01'),
(2, 2, '2019-02-25', 2500000, 9500, '2019-05-25 00:00:00', '2019-02-26 15:25:39', '2019-02-28 12:36:47'),
(3, 3, '2019-02-25', 45000, 0, NULL, '2019-02-28 17:18:50', '2019-02-28 17:18:50'),
(5, 12, '2019-03-25', 45000, 0, NULL, '2019-03-25 07:08:43', '2019-03-25 07:08:43'),
(7, 12, '2019-03-25', 20000, 0, NULL, '2019-03-25 07:13:13', '2019-03-25 07:13:13');

-- --------------------------------------------------------

--
-- Table structure for table `tax_status`
--

CREATE TABLE `tax_status` (
  `id` int(11) NOT NULL,
  `status` varchar(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Dumping data for table `tax_status`
--

INSERT INTO `tax_status` (`id`, `status`) VALUES
(1, 'paid'),
(2, 'pending');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `admins`
--
ALTER TABLE `admins`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `id_UNIQUE` (`id`),
  ADD UNIQUE KEY `token_UNIQUE` (`token`),
  ADD UNIQUE KEY `email_UNIQUE` (`email`);

--
-- Indexes for table `businesses`
--
ALTER TABLE `businesses`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `id_UNIQUE` (`id`),
  ADD UNIQUE KEY `email_UNIQUE` (`email`);

--
-- Indexes for table `messages`
--
ALTER TABLE `messages`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `id_UNIQUE` (`id`);

--
-- Indexes for table `notifications`
--
ALTER TABLE `notifications`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `id_UNIQUE` (`id`),
  ADD KEY `fk_notifications_1_idx` (`business_id`);

--
-- Indexes for table `tax`
--
ALTER TABLE `tax`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `id_UNIQUE` (`id`),
  ADD KEY `fk_tax_1_idx` (`business_id`);

--
-- Indexes for table `tax_status`
--
ALTER TABLE `tax_status`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `status_UNIQUE` (`status`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `admins`
--
ALTER TABLE `admins`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=7;
--
-- AUTO_INCREMENT for table `businesses`
--
ALTER TABLE `businesses`
  MODIFY `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=13;
--
-- AUTO_INCREMENT for table `messages`
--
ALTER TABLE `messages`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;
--
-- AUTO_INCREMENT for table `notifications`
--
ALTER TABLE `notifications`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT;
--
-- AUTO_INCREMENT for table `tax`
--
ALTER TABLE `tax`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=8;
--
-- Constraints for dumped tables
--

--
-- Constraints for table `notifications`
--
ALTER TABLE `notifications`
  ADD CONSTRAINT `fk_notifications_1` FOREIGN KEY (`business_id`) REFERENCES `businesses` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION;

--
-- Constraints for table `tax`
--
ALTER TABLE `tax`
  ADD CONSTRAINT `fk_tax_1` FOREIGN KEY (`business_id`) REFERENCES `businesses` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
