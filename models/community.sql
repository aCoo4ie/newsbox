DROP TABLE IF EXISTS `community`;
-- Table structure for table `community`
CREATE TABLE community (
    id INT PRIMARY KEY AUTO_INCREMENT,
    community_id INT UNIQUE NOT NULL,
    community_name VARCHAR(128) NOT NULL,
    introduction VARCHAR(256) NOT NULL,
    create_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    update_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- Records of `community`
INSERT INTO community (community_id, community_name, introduction) VALUES
(1, 'Tech Frontiers', 'Discuss the latest technological trends and innovative technologies.'),
(2, 'Programming Technology', 'Share programming tips, experiences, and the latest development tools.'),
(3, 'Movie Reviews', 'Share and discuss the latest movies, reviews, and filmmaking.'),
(4, 'Food Sharing', 'Share food recipes, restaurant recommendations, and cooking tips.'),
(5, 'Travel Guide', 'Share travel experiences, attraction recommendations, and travel tips.'),
(6, 'Game World', 'Discuss various games, share game strategies and experiences.'),
(7, 'Sports Events', 'Discuss the latest sports events and athletes.'),
(8, 'Music Appreciation', 'Share musical works, musicians, and music culture.'),
(9, 'Literature Exchange', 'Share literary works, reading experiences, and literary reviews.'),
(10, 'Career Development', 'Share workplace experiences, job search skills, and career development advice.');