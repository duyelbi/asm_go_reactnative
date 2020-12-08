-- phpMyAdmin SQL Dump
-- version 5.0.4
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Dec 06, 2020 at 08:50 AM
-- Server version: 10.4.16-MariaDB
-- PHP Version: 7.4.12

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `quizdb`
--

-- --------------------------------------------------------

--
-- Table structure for table `answers`
--

CREATE TABLE `answers` (
  `answer_id` int(11) NOT NULL,
  `answer` varchar(200) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `answers`
--

INSERT INTO `answers` (`answer_id`, `answer`) VALUES
(11, '2'),
(12, '3'),
(13, '4'),
(14, '5'),
(21, '2'),
(22, '3'),
(23, '4'),
(24, '5'),
(31, '4'),
(32, '5'),
(33, '6'),
(34, '7'),
(41, '55'),
(42, '56'),
(43, '57'),
(44, '58'),
(51, '19'),
(52, '20'),
(53, '19'),
(54, '20'),
(61, '20'),
(62, '21'),
(63, '22'),
(64, '23'),
(71, '2'),
(72, '3'),
(73, '4'),
(74, '5'),
(81, '2'),
(82, '3'),
(83, '4'),
(84, '5'),
(91, '36'),
(92, '37'),
(93, '38'),
(94, '39'),
(101, '43'),
(102, '42'),
(103, '41'),
(104, '40'),
(111, '77'),
(112, '78'),
(113, '79'),
(114, '80'),
(121, '19'),
(122, '18'),
(123, '17'),
(124, '16'),
(131, '1'),
(132, '2'),
(133, '3'),
(134, '4'),
(141, '1'),
(142, '2'),
(143, '3'),
(144, '4'),
(151, '1'),
(152, '2'),
(153, '3'),
(154, '4'),
(161, '1'),
(162, '2'),
(163, '3'),
(164, '4'),
(171, '1'),
(172, '2'),
(173, '3'),
(174, '4'),
(181, '1'),
(182, '2'),
(183, '3'),
(184, '4'),
(191, '1'),
(192, '2'),
(193, '3'),
(194, '4'),
(201, '1'),
(202, '2'),
(203, '3'),
(204, '4'),
(211, '1'),
(212, '2'),
(213, '3'),
(214, '4'),
(221, '1'),
(222, '2'),
(223, '3'),
(224, '4'),
(231, '1'),
(232, '2'),
(233, '3'),
(234, '4');

-- --------------------------------------------------------

--
-- Table structure for table `questions`
--

CREATE TABLE `questions` (
  `question_id` int(11) NOT NULL,
  `difficulty` varchar(200) DEFAULT NULL,
  `question` varchar(200) DEFAULT NULL,
  `correct_answer` varchar(200) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `questions`
--

INSERT INTO `questions` (`question_id`, `difficulty`, `question`, `correct_answer`) VALUES
(1, 'easy', '1 + 1', '2'),
(2, 'easy', '1 + 2', '3'),
(3, 'medium', '2 x 3', '6'),
(4, 'medium', '7 X 8', '56'),
(5, 'hard', '5 X 5 - 6', '19'),
(6, 'hard', '9 + 4 x 3', '21'),
(7, 'easy', '3 - 1', '2'),
(8, 'easy', '9 - 4', '5'),
(9, 'mideum', '4 x 9', '36'),
(10, 'easy', '5 x 8', '40'),
(11, 'hard', '9  x 8 + 5', '77'),
(12, 'hard', '98/2 -32', '17');

-- --------------------------------------------------------

--
-- Table structure for table `question_answer`
--

CREATE TABLE `question_answer` (
  `question_id` int(11) NOT NULL,
  `answer_id` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `question_answer`
--

INSERT INTO `question_answer` (`question_id`, `answer_id`) VALUES
(1, 11),
(1, 12),
(1, 13),
(1, 14),
(2, 21),
(2, 22),
(2, 23),
(2, 24),
(3, 31),
(3, 32),
(3, 33),
(3, 34),
(4, 41),
(4, 42),
(4, 43),
(4, 44),
(5, 51),
(5, 52),
(5, 53),
(5, 54),
(6, 61),
(6, 62),
(6, 63),
(6, 64),
(7, 71),
(7, 72),
(7, 73),
(7, 74),
(8, 81),
(8, 82),
(8, 83),
(8, 84),
(9, 91),
(9, 92),
(9, 93),
(9, 94),
(10, 101),
(10, 102),
(10, 103),
(10, 144),
(11, 111),
(11, 112),
(11, 113),
(11, 114),
(12, 121),
(12, 122),
(12, 123),
(12, 124);

--
-- Indexes for dumped tables
--

--
-- Indexes for table `answers`
--
ALTER TABLE `answers`
  ADD PRIMARY KEY (`answer_id`);

--
-- Indexes for table `questions`
--
ALTER TABLE `questions`
  ADD PRIMARY KEY (`question_id`);

--
-- Indexes for table `question_answer`
--
ALTER TABLE `question_answer`
  ADD PRIMARY KEY (`question_id`,`answer_id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `answers`
--
ALTER TABLE `answers`
  MODIFY `answer_id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=235;

--
-- AUTO_INCREMENT for table `questions`
--
ALTER TABLE `questions`
  MODIFY `question_id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=13;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;

-- seclect all questions and answers in db
SELECT q.*, a.* FROM questions q INNER JOIN question_answer qa ON qa.question_id = q.question_id INNER JOIN answers a ON a.answer_id = qa.answer_id