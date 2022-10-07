-- phpMyAdmin SQL Dump
-- version 5.2.0
-- https://www.phpmyadmin.net/
--
-- Máy chủ: 127.0.0.1
-- Thời gian đã tạo: Th10 07, 2022 lúc 10:28 AM
-- Phiên bản máy phục vụ: 10.4.25-MariaDB
-- Phiên bản PHP: 8.1.10

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Cơ sở dữ liệu: `developer`
--

-- --------------------------------------------------------

--
-- Cấu trúc bảng cho bảng `character_example`
--

CREATE TABLE `character_example` (
  `id` int(11) NOT NULL,
  `simplified` varchar(1500) NOT NULL,
  `hanzi` varchar(1500) NOT NULL,
  `pinyin` varchar(1500) NOT NULL,
  `translation` varchar(1500) NOT NULL,
  `audio` varchar(1500) NOT NULL,
  `level` varchar(1500) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Đang đổ dữ liệu cho bảng `character_example`
--

INSERT INTO `character_example` (`id`, `simplified`, `hanzi`, `pinyin`, `translation`, `audio`, `level`, `created_at`, `updated_at`) VALUES
(33, '固', '这里写的维修费用应该注明是屋内装修，固定的东西，房客自己买的家电得自己修。', 'zhèlǐ xiě de wéixiū fèiyòng yīnggāi zhùmíng shì wū nèi zhuāngxiū ，gùdìng de dōngxi ，fángkè zìjǐ mǎi de jiā diàn děi zìjǐ xiū 。', 'Here it specifies that the maintenance costs apply to interior decoration and fixed appliances, and that the tenant has to fix consumer electronics by themself.', 'https://s3contents.chinesepod.com/4365/4883c6d7050059613d320dac49037d8c0e4cdf47/mp3/64/chinesepod_8330_dialogue_79816_prototype_1597713530.mp3', 'Intermediate', '2022-10-04 08:55:53', '2022-10-04 08:55:53'),
(34, '固', '因此水星只是没有液态水，水是以固态冰的形式存在。', 'yīncǐ Shuǐxīng zhǐshì méiyǒu yètàishuǐ ，shuǐ shì yǐ gùtài bīng de xíngshì cúnzài 。', 'Therefore, there is no liquid water on Mercury, but it is existed in the form of solid ice.', 'https://s3contents.chinesepod.com/4046/b3d7652da4ee7c8ce80b81cf8b200356f77d6b6f/mp3/64/chinesepod_7843_dialogue_74831_prototype_1509508531.mp3', 'Advanced', '2022-10-04 08:55:53', '2022-10-04 08:55:53'),
(35, '固', '你们真是年轻有为，不过，别怪我老生常谈，实现自我价值固然重要，也别忽略了身体健康，毕竟身体是革命的本钱呀！', 'nǐmen zhēn shì niánqīngyǒuwéi ，bùguò ，bié guài wǒ lǎoshēng cháng tán ，shíxiàn zìwǒ jiàzhí gùrán zhòngyào ，yě bié hūlüè le shēntǐ jiànkāng ，bìjìng shēntǐ shì gémìng de běnqián ya ！', 'You really are a young up-and-comer, but don\'t hold against me if I spout some cliches. Self realization really is important, and don\'t ignore the value of your health, one\'s body is one\'s best asset in a revolution!', 'https://s3contents.chinesepod.com/3039/c673d2850160104fc18d16ed35c1fa70ed5dc14d/mp3/64/chinesepod_7708_dialogue_73150_prototype_52148_1856_8348_57233_70634_45871.mp3', 'Advanced', '2022-10-04 08:55:53', '2022-10-04 08:55:53'),
(36, '固', '说到契约，记住，写得越详细越好，建材的项目、规格、数量、品牌等等都要写清楚喔，还有，千万不要忘记要求建商提供保固，虽然贵一点，但是比较有保障。', 'shuōdào qìyuē ，jìzhu ，xiě de yuè xiángxì yuè hǎo ，jiàncái de xiàngmù 、guīgé 、shùliàng 、pǐnpái děngdeng dōuyào xiě qīngchu ō ，háiyǒu ，qiānwàn bù yào wàngjì yāoqiú jiànshāng tígōng bǎogù ，suīrán guì yīdiǎn ，dànshì bǐjiào yǒu bǎozhàng 。', 'Talking of contracts, remember, the more detailed you write them the better and specify clearly things including the buildings materials, the specifications, the amounts, the brands and all that stuff. As well as that, whatever you do, don\'t forget to request that the construction company provides a warranty. Although it\'s a little expensive, it provides you with more of a safeguard.', 'https://s3contents.chinesepod.com/2978/5b62cc28b92af37d7a4ef181d1db482cc163324b/mp3/64/chinesepod_7633_dialogue_72232_prototype_41630_89842_114022_41818_114349_14252_9788_140618.mp3', 'Upper Intermediate', '2022-10-04 08:55:53', '2022-10-04 08:55:53'),
(37, '固', '我也知道现在这时代不应该再固守成规、遵循“按资排辈”那一套了，只是被一个年轻人牵着鼻子走，心理多少有些不平衡。', 'wǒ yě zhīdào xiànzài zhè shídài bù yīnggāi zài gùshǒu chéngguī 、zūnxún “ànzīpáibèi ”nà yī tào le ，zhǐshì bèi yī ge niánqīngrén qiān zhe bízi zǒu ，xīnlǐ duōshǎo yǒuxiē bù pínghéng 。', 'I am aware in today’s times, we shouldn’t stick to rules where positions are based on seniority. But I still feel mentally awkward to be led by someone younger.', 'https://s3contents.chinesepod.com/2921/21a864195076bc41bd7a47851d222c62586c84fa/mp3/64/chinesepod_7569_dialogue_71474_prototype_21311_86088_115482_105660.mp3', 'Upper Intermediate', '2022-10-04 08:55:53', '2022-10-04 08:55:53'),
(38, '固', '死鸭子嘴硬”表示一个人很固执。大家都觉得他是错的，他却不承认自己是错的。', 'sǐyāzizuǐyìng ”biǎoshì yī ge rén hěn gùzhí 。dàjiā dōu juéde tā shì cuò de ，tā què bù chéngrèn zìjǐ shì cuò de 。', '“A dead duck’s mouth is hard.” means the person is stubborn to a fault. Everyone thinks he is wrong, but he doesn’t admit it.', 'https://s3contents.chinesepod.com/2881/5f5ed0d965da3cc0204fd48fda48450167c6edc9/mp3/64/chinesepod_7520_dialogue_70993_prototype_160132_180832.mp3', 'Intermediate', '2022-10-04 08:55:53', '2022-10-04 08:55:53'),
(39, '固', '对不听劝阻的吸烟者，要求其离开该场所；对不听劝阻且不离开该场所的，应当固定相关证据并向有关行政管理部门举报；对不听劝阻并扰乱公共秩序的，向公安机关报案。', 'duì bù tīng quànzǔ de xīyānzhě ，yāoqiú qí líkāi gāi chǎngsuǒ ；duì bù tīng quànzǔ qiě bù líkāi gāi chǎngsuǒ de ，yīngdāng gùdìng xiāngguān zhèngjù bìng xiàng yǒuguān xíngzhèng guǎnlǐ bùmén jǔbào ；duì bù tīng quànzǔ bìng rǎoluàn gōnggòng zhìxù de ，xiàng gōngānjīguān bàoàn 。', 'Offenders who do not obey will be asked to leave and those who refuse to obey and leave will be reported to the relevant administrative department with the necessary evidence. Finally offenders who refuse to listen and cause public disorder will be reported to public security.', 'https://s3contents.chinesepod.com/4149/c8bb3d82d9110dc5ab18ee87ba661c3ce97b4589/mp3/64/chinesepod_8065_dialogue_77082_prototype_1545797414.mp3', 'Media', '2022-10-04 08:55:53', '2022-10-04 08:55:53'),
(40, '固', '为了和平，我们要牢固树立人类命运共同体意识。偏见和歧视、仇恨和战争，只会带来灾难和痛苦。相互尊重、平等相处、和平发展、共同繁荣，才是人间正道，世界各国应该共同维护以联合国宪章宗旨和原则为核心的国际秩序和国际体系，积极构建以合作共赢为核心的新型国际关系共同推进世界和平与发展的崇高事业。', 'wèile hépíng ，wǒmen yào láogù shùlì rénlèi mìngyùn gòngtóngtǐ yìshi 。piānjiàn hé qíshì 、chóuhèn hé zhànzhēng ，zhī huì dàilái zāinàn hé tòngkǔ 。xiānghù zūnzhòng 、píngděng xiāngchǔ 、hépíng fāzhǎn 、gòngtóng fánróng ，cái shì rénjiān zhèngdào ，shìjiègèguó yīnggāi gòngtóng wéihù yǐ Liánhéguó xiànzhāng zōngzhǐ hé yuánzé wèi héxīn de guójì zhìxù hé guójì tǐxì ，jījí gòujiàn yǐ to collaborate gòngyíng wèi héxīn de xīnxíng guójì guānxi gòngtóng tuījìn shìjiè hépíng yǔ fāzhǎn de chónggāo shìyè 。', 'For peace, we have to securely establish an awareness of a community of destiny among humanity. Prejudice and discrimination, hostility and war, they only bring disaster and pain. Mutual respect, treating one another as equals, peaceful development, prospering together, this is the true path for humanity. Each country in the world should uphold the objectives and principles enshrined in the United Nations charter as the core tenets of international order and the international system, to proactively construct new international relations with win-win cooperation at its core, to continue to carry forward the sublime undertaking that is world peace and development.', 'https://s3contents.chinesepod.com/4300/edf19466fe8131b4990f62f3dd367f945ec2b90c/mp3/64/chinesepod_8253_dialogue_79192_prototype_1586146485.mp3', 'Advanced', '2022-10-04 08:55:53', '2022-10-04 08:55:53'),
(41, '固', '贵公司专营加拿大、美国和澳洲的留学业务，特别专攻社会科学类的学门，不过你们在加拿大的市场特别大、也做得很用心，反而对于竞争激烈的美国只针对固定的五、六所顶尖学院，其它知名大学却没有花什么工夫去经营，所以我觉得我的加拿大背景很有优势。', 'guì gōngsī zhuānyíng Jiānádà 、Měiguó hé Àozhōu de liúxué yèwù ，tèbié zhuāngōng shèhuìkēxué lèi de xuémén ，bùguò nǐmen zài Jiānádà de shìchǎng tèbié dà 、yě zuò de hěn yòngxīn ，fǎnér duìyú jìngzhēng jīliè de Měiguó zhī zhēnduì gùdìng de wǔ 、liù suǒ dǐngjiān xuéyuàn ，qítā zhīmíng dàxué què méiyǒu huā shénme gōngfū qù jīngyíng ，suǒyǐ wǒ juéde wǒ de Jiānádà bèijǐng hěnyǒu yōushì 。', 'Your company specializes in facilitating exchange programs in Canada, the US and Australia, especially for the social sciences. But you have a particularly strong market in Canada, and you are very meticulous there. In the more competitive US market, you only target five or six of the top schools. You don\'t spend as much energy on other famous universities, so I think my background in Canada is a big strength.', 'https://s3contents.chinesepod.com/4150/ca48f4b02d3bffe83851b87eaec16152691407f6/mp3/64/chinesepod_8057_dialogue_76980_prototype_1542882942.mp3', 'Advanced', '2022-10-04 08:55:53', '2022-10-04 08:55:53'),
(42, '固', '没错，科技的力量固然推动了社会发展，可是我多么希望可以回到原来单纯美好的朋友圈呀！', 'méicuò ，kējì de lìliang gùrán tuīdòng le shèhuì fāzhǎn ，kěshì wǒ duōme xīwàng kěyǐ huí dào yuánlái dānchún měihǎo de péngyou quān ya ！', 'That\'s right, the power of technology does of course drive social development, but I really wish to return to the pure and simple old days of Moments.', 'https://s3contents.chinesepod.com/3013/5313e0aaf78c0ba0040702c3d59c43ceb648ca85/mp3/64/chinesepod_7663_dialogue_72601_prototype_23655_71466.mp3', 'Upper Intermediate', '2022-10-04 08:55:53', '2022-10-04 08:55:53');

-- --------------------------------------------------------

--
-- Cấu trúc bảng cho bảng `character_lookup`
--

CREATE TABLE `character_lookup` (
  `id` int(11) NOT NULL,
  `simplified` varchar(150) NOT NULL,
  `rank` int(11) NOT NULL,
  `hsk` int(11) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Đang đổ dữ liệu cho bảng `character_lookup`
--

INSERT INTO `character_lookup` (`id`, `simplified`, `rank`, `hsk`, `created_at`, `updated_at`) VALUES
(1, '是', 4, 1, '2022-10-04 08:17:33', '2022-10-04 08:17:33'),
(2, '老', 294, 3, '2022-10-04 08:18:20', '2022-10-04 08:18:20'),
(3, '采', 3558, 0, '2022-10-04 08:18:57', '2022-10-04 08:18:57'),
(4, '产', 1543, 0, '2022-10-04 08:28:27', '2022-10-04 08:28:27'),
(5, '访', 3145, 0, '2022-10-04 08:28:33', '2022-10-04 08:28:33'),
(6, '汇', 2699, 0, '2022-10-04 08:28:38', '2022-10-04 08:28:38'),
(7, '缓', 3996, 0, '2022-10-04 08:28:40', '2022-10-04 08:28:40'),
(8, '固', 5882, 0, '2022-10-04 08:55:51', '2022-10-04 08:55:51');

-- --------------------------------------------------------

--
-- Cấu trúc bảng cho bảng `entry`
--

CREATE TABLE `entry` (
  `id` int(11) NOT NULL,
  `simplified` varchar(150) NOT NULL,
  `traditional` varchar(150) NOT NULL,
  `pinyin` varchar(150) NOT NULL,
  `definitions` varchar(2000) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Đang đổ dữ liệu cho bảng `entry`
--

INSERT INTO `entry` (`id`, `simplified`, `traditional`, `pinyin`, `definitions`, `created_at`, `updated_at`) VALUES
(1, '是', '是', 'shì', '[\"is\",\"are\",\"am\",\"yes\",\"to be\"]', '2022-10-04 08:17:33', '2022-10-04 08:17:33'),
(2, '老', '老', 'lǎo', '[\"prefix used before the surname of a person or a numeral indicating the order of birth of the children in a family or to indicate affection or familiarity\",\"old (of people)\",\"venerable (person)\",\"experienced\",\"of long standing\",\"always\",\"all the time\",\"of the past\",\"very\",\"outdated\",\"(of meat etc) tough\"]', '2022-10-04 08:18:20', '2022-10-04 08:18:20'),
(3, '采', '採', 'cǎi', '[\"to pick\",\"to pluck\",\"to collect\",\"to select\",\"to choose\",\"to gather\"]', '2022-10-04 08:18:58', '2022-10-04 08:18:58'),
(4, '采', '采', 'cǎi', '[\"color\",\"complexion\",\"looks\",\"variant of 彩[cai3]\",\"variant of 採|采[cai3]\"]', '2022-10-04 08:18:58', '2022-10-04 08:18:58'),
(5, '采', '采', 'cài', '[\"allotment to a feudal noble\"]', '2022-10-04 08:18:58', '2022-10-04 08:18:58'),
(6, '产', '產', 'chǎn', '[\"to give birth\",\"to reproduce\",\"to produce\",\"product\",\"resource\",\"estate\",\"property\"]', '2022-10-04 08:28:27', '2022-10-04 08:28:27'),
(7, '访', '訪', 'fǎng', '[\"to visit\",\"to call on\",\"to seek\",\"to inquire\",\"to investigate\"]', '2022-10-04 08:28:33', '2022-10-04 08:28:33'),
(8, '汇', '匯', 'huì', '[\"to remit\",\"to converge (of rivers)\",\"to exchange\"]', '2022-10-04 08:28:38', '2022-10-04 08:28:38'),
(9, '汇', '彙', 'huì', '[\"class\",\"collection\"]', '2022-10-04 08:28:38', '2022-10-04 08:28:38'),
(10, '缓', '緩', 'huǎn', '[\"slow\",\"unhurried\",\"sluggish\",\"gradual\",\"not tense\",\"relaxed\",\"to postpone\",\"to defer\",\"to stall\",\"to stave off\",\"to revive\",\"to recuperate\"]', '2022-10-04 08:28:40', '2022-10-04 08:28:40'),
(11, '固', '固', 'gù', '[\"hard\",\"strong\",\"solid\",\"sure\",\"assuredly\",\"undoubtedly\",\"of course\",\"indeed\",\"admittedly\"]', '2022-10-04 08:55:51', '2022-10-04 08:55:51');

-- --------------------------------------------------------

--
-- Cấu trúc bảng cho bảng `user`
--

CREATE TABLE `user` (
  `id` int(11) NOT NULL,
  `email` varchar(150) NOT NULL,
  `password` varchar(150) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `password_updated_at` timestamp NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Đang đổ dữ liệu cho bảng `user`
--

INSERT INTO `user` (`id`, `email`, `password`, `created_at`, `password_updated_at`) VALUES
(18, 'truaong@k0test.com', '8d969eef6ecad3c29a3a629280e686cf0c3f5d5a86aff3ca12020c923adc6c92', '2022-10-06 10:04:24', '2022-10-06 10:04:24'),
(19, 'truong@test.com', '8d969eef6ecad3c29a3a629280e686cf0c3f5d5a86aff3ca12020c923adc6c92', '2022-10-06 10:04:57', '2022-10-07 08:27:19');

--
-- Chỉ mục cho các bảng đã đổ
--

--
-- Chỉ mục cho bảng `character_example`
--
ALTER TABLE `character_example`
  ADD PRIMARY KEY (`id`);

--
-- Chỉ mục cho bảng `character_lookup`
--
ALTER TABLE `character_lookup`
  ADD PRIMARY KEY (`id`);

--
-- Chỉ mục cho bảng `entry`
--
ALTER TABLE `entry`
  ADD PRIMARY KEY (`id`);

--
-- Chỉ mục cho bảng `user`
--
ALTER TABLE `user`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT cho các bảng đã đổ
--

--
-- AUTO_INCREMENT cho bảng `character_example`
--
ALTER TABLE `character_example`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=43;

--
-- AUTO_INCREMENT cho bảng `character_lookup`
--
ALTER TABLE `character_lookup`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=9;

--
-- AUTO_INCREMENT cho bảng `entry`
--
ALTER TABLE `entry`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=12;

--
-- AUTO_INCREMENT cho bảng `user`
--
ALTER TABLE `user`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=20;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
