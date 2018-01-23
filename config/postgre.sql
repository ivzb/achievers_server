-- *****************************************************************************
-- // Clean up 
-- ****************************************************************************/
DROP FUNCTION IF EXISTS set_updated_at;

DROP DATABASE IF EXISTS achievers;

-- *****************************************************************************
-- // Create new database
-- ****************************************************************************/
CREATE DATABASE "achievers"
    WITH OWNER "root"
    ENCODING 'UTF8'
    LC_COLLATE = 'en_US.UTF-8'
    LC_CTYPE = 'en_US.UTF-8';

-- *****************************************************************************
-- // Attach to db
-- ****************************************************************************/
\c achievers;

-- *****************************************************************************
-- // Required extensions
-- ****************************************************************************/
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";


-- *****************************************************************************
-- // Update timestamp trigger
-- ****************************************************************************/
CREATE FUNCTION set_updated_at()
  RETURNS TRIGGER
  LANGUAGE plpgsql
AS $$
BEGIN
  NEW.updated_at := now();
  RETURN NEW;
END;
$$;

-- *****************************************************************************
-- // user_status
-- ****************************************************************************/
CREATE TABLE user_status (
    id smallserial NOT NULL,
    
    status VARCHAR(25) NOT NULL,
    
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NOT NULL DEFAULT '0001-01-01 00:00:00',
    
    PRIMARY KEY (id)
);

CREATE TRIGGER user_status_updated_at
  BEFORE UPDATE ON user_status
  FOR EACH ROW EXECUTE PROCEDURE set_updated_at();

-- *****************************************************************************
-- // involvement 
-- ****************************************************************************/

CREATE TABLE involvement (
    id smallserial NOT NULL,
    
    title VARCHAR(25) NOT NULL,
    
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NOT NULL DEFAULT '0001-01-01 00:00:00',
    
    PRIMARY KEY (id)
);

CREATE TRIGGER involvement_updated_at
  BEFORE UPDATE ON involvement
  FOR EACH ROW EXECUTE PROCEDURE set_updated_at();

-- *****************************************************************************
-- // multimedia_type
-- ****************************************************************************/

CREATE TABLE multimedia_type (
    id smallserial NOT NULL,
    
    title VARCHAR(25) NOT NULL,
    
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NOT NULL DEFAULT '0001-01-01 00:00:00',
    
    PRIMARY KEY (id)
);

CREATE TRIGGER multimedia_type_updated_at
  BEFORE UPDATE ON multimedia_type
  FOR EACH ROW EXECUTE PROCEDURE set_updated_at();

-- *****************************************************************************
-- // reward_type
-- ****************************************************************************/

CREATE TABLE reward_type (
    id smallserial NOT NULL,
    
    title VARCHAR(25) NOT NULL,
    
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NOT NULL DEFAULT '0001-01-01 00:00:00',
    
    PRIMARY KEY (id)
);

CREATE TRIGGER reward_type_updated_at
  BEFORE UPDATE ON reward_type
  FOR EACH ROW EXECUTE PROCEDURE set_updated_at();

-- *****************************************************************************
-- // quest_type
-- ****************************************************************************/

CREATE TABLE quest_type (
    id smallserial NOT NULL,
    
    title VARCHAR(25) NOT NULL,
    
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NOT NULL DEFAULT '0001-01-01 00:00:00',
    
    PRIMARY KEY (id)
);

CREATE TRIGGER quest_type_updated_at
  BEFORE UPDATE ON quest_type
  FOR EACH ROW EXECUTE PROCEDURE set_updated_at();

-- *****************************************************************************
-- // user
-- ****************************************************************************/

CREATE TABLE "user" (
    id UUID NOT NULL DEFAULT uuid_generate_v4(),
    
    email      VARCHAR(100) NOT NULL,
    password   BYTEA NOT NULL,
    
    user_status_id smallint NOT NULL DEFAULT 1,
    
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NOT NULL DEFAULT '0001-01-01 00:00:00',
    
    UNIQUE (email),
    CONSTRAINT f_user__user_status FOREIGN KEY (user_status_id)
        REFERENCES user_status (id) ON DELETE NO ACTION ON UPDATE NO ACTION,
    
    PRIMARY KEY (id)
);

CREATE TRIGGER user_updated_at
  BEFORE UPDATE ON "user"
  FOR EACH ROW EXECUTE PROCEDURE set_updated_at();

-- *****************************************************************************
-- // profile
-- ****************************************************************************/
CREATE TABLE profile (
    id UUID NOT NULL DEFAULT uuid_generate_v4(),

    name VARCHAR(255) NOT NULL,

    user_id UUID NOT NULL,

    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NOT NULL DEFAULT '0001-01-01 00:00:00',

    CONSTRAINT f_profile__user FOREIGN KEY (user_id) 
        REFERENCES "user" (id) ON DELETE NO ACTION ON UPDATE NO ACTION,
    
    PRIMARY KEY (id)
);

CREATE TRIGGER profile_updated_at
  BEFORE UPDATE ON profile
  FOR EACH ROW EXECUTE PROCEDURE set_updated_at();

-- *****************************************************************************
-- // achievement 
-- ****************************************************************************/
CREATE TABLE achievement (
    id UUID NOT NULL DEFAULT uuid_generate_v4(),
    
    title       VARCHAR(50) NOT NULL,
    description VARCHAR(255) NOT NULL,
    picture_url VARCHAR(100) NOT NULL,
    
    involvement_id smallint NOT NULL,
    user_id        UUID NOT NULL,

    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NOT NULL DEFAULT '0001-01-01 00:00:00',

    CONSTRAINT f_achievement__involvement FOREIGN KEY (involvement_id)
        REFERENCES involvement (id) ON DELETE NO ACTION ON UPDATE NO ACTION,
    
    CONSTRAINT f_achievement__user FOREIGN KEY (user_id) 
        REFERENCES "user" (id) ON DELETE NO ACTION ON UPDATE NO ACTION,
    
    PRIMARY KEY (id)
);

CREATE TRIGGER achievement_updated_at
  BEFORE UPDATE ON achievement 
  FOR EACH ROW EXECUTE PROCEDURE set_updated_at();

-- *****************************************************************************
-- // evidence 
-- ****************************************************************************/
CREATE TABLE evidence (
    id UUID NOT NULL DEFAULT uuid_generate_v4(),
    
    title       VARCHAR(255) NOT NULL,
    picture_url VARCHAR(255) NOT NULL,
    url         VARCHAR(255) NOT NULL,
    
    multimedia_type_id smallint NOT NULL,
    achievement_id     UUID NOT NULL,
    user_id            UUID NOT NULL,

    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NOT NULL DEFAULT '0001-01-01 00:00:00',

    CONSTRAINT f_evidence__multimedia_type FOREIGN KEY (multimedia_type_id)
        REFERENCES multimedia_type (id) ON DELETE NO ACTION ON UPDATE NO ACTION,

    CONSTRAINT f_evidence__achievement FOREIGN KEY (achievement_id) 
        REFERENCES achievement (id) ON DELETE NO ACTION ON UPDATE NO ACTION,
    
    CONSTRAINT f_evidence__user FOREIGN KEY (user_id) 
        REFERENCES "user" (id) ON DELETE NO ACTION ON UPDATE NO ACTION,

    PRIMARY KEY (id)
);

CREATE TRIGGER evidence_updated_at
  BEFORE UPDATE ON evidence 
  FOR EACH ROW EXECUTE PROCEDURE set_updated_at();

-- *****************************************************************************
-- // reward
-- ****************************************************************************/
CREATE TABLE reward (
    id UUID NOT NULL DEFAULT uuid_generate_v4(),
    
    title       VARCHAR(255) NOT NULL,
    description VARCHAR(255) NOT NULL,
    picture_url VARCHAR(255) NOT NULL,
    
    reward_type_id smallint NOT NULL,
    user_id            UUID NOT NULL,

    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NOT NULL DEFAULT '0001-01-01 00:00:00',

    CONSTRAINT f_reward__reward_type FOREIGN KEY (reward_type_id)
        REFERENCES reward_type (id) ON DELETE NO ACTION ON UPDATE NO ACTION,

    CONSTRAINT f_reward__user FOREIGN KEY (user_id) 
        REFERENCES "user" (id) ON DELETE NO ACTION ON UPDATE NO ACTION,
    
    PRIMARY KEY (id)
);

CREATE TRIGGER reward_updated_at
  BEFORE UPDATE ON reward
  FOR EACH ROW EXECUTE PROCEDURE set_updated_at();

-- *****************************************************************************
-- // quest
-- ****************************************************************************/
CREATE TABLE quest (
    id UUID NOT NULL DEFAULT uuid_generate_v4(),
    
    title       VARCHAR(255) NOT NULL,
    picture_url VARCHAR(255) NOT NULL,
    
    involvement_id smallint NOT NULL,
    quest_type_id  smallint NOT NULL,
    user_id        UUID NOT NULL,

    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NOT NULL DEFAULT '0001-01-01 00:00:00',

    CONSTRAINT f_quest__involvement FOREIGN KEY (involvement_id)
        REFERENCES involvement (id) ON DELETE NO ACTION ON UPDATE NO ACTION,

    CONSTRAINT f_quest__quest_type FOREIGN KEY (quest_type_id)
        REFERENCES quest_type (id) ON DELETE NO ACTION ON UPDATE NO ACTION,

    CONSTRAINT f_quest__user FOREIGN KEY (user_id) 
        REFERENCES "user" (id) ON DELETE NO ACTION ON UPDATE NO ACTION,

    PRIMARY KEY (id)
);

CREATE TRIGGER quest_updated_at
  BEFORE UPDATE ON quest 
  FOR EACH ROW EXECUTE PROCEDURE set_updated_at();

-- *****************************************************************************
-- // quest_achievement
-- ****************************************************************************/
CREATE TABLE quest_achievement (
    id UUID NOT NULL DEFAULT uuid_generate_v4(),
    
    quest_id       UUID NOT NULL,
    achievement_id UUID NOT NULL,
    user_id        UUID NOT NULL,

    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NOT NULL DEFAULT '0001-01-01 00:00:00',

    CONSTRAINT f_quest_achievement__quest FOREIGN KEY (quest_id)
        REFERENCES quest (id) ON DELETE NO ACTION ON UPDATE NO ACTION,

    CONSTRAINT f_quest_achievement__achievement FOREIGN KEY (achievement_id)
        REFERENCES achievement (id) ON DELETE NO ACTION ON UPDATE NO ACTION,

    CONSTRAINT f_quest_achievement__user FOREIGN KEY (user_id) 
        REFERENCES "user" (id) ON DELETE NO ACTION ON UPDATE NO ACTION,

    PRIMARY KEY (id)
);

CREATE TRIGGER quest_achievement_updated_at
  BEFORE UPDATE ON quest_achievement 
  FOR EACH ROW EXECUTE PROCEDURE set_updated_at();

/* *****************************************************************************
// Seed tables
// ****************************************************************************/
INSERT INTO user_status (id, status) VALUES
(1, 'active'),
(2, 'inactive');

INSERT INTO involvement (id, title) VALUES
(1, 'bronze'),
(2, 'silver'),
(3, 'gold'),
(4, 'platinum'),
(5, 'diamond');

INSERT INTO multimedia_type (id, title) VALUES
(1, 'photo'),
(2, 'video'),
(3, 'voice');

INSERT INTO reward_type (id, title) VALUES
(1, 'experience'),
(2, 'item'),
(3, 'title');

INSERT INTO quest_type (id, title) VALUES
(1, 'world'),
(2, 'daily'),
(3, 'weekly'),
(4, 'monthly');
