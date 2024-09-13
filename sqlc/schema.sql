CREATE TABLE IF NOT EXISTS "users" (
                                       "id" bigint NOT NULL UNIQUE,
                                       "user_role_id" bigint NOT NULL,
                                       "name" varchar(255) NOT NULL,
    "password" varchar(255) NOT NULL,
    PRIMARY KEY ("id")
    );

CREATE TABLE IF NOT EXISTS "comments" (
                                          "id" bigint NOT NULL,
                                          "user_id" bigint NOT NULL,
                                          "news_id" bigint NOT NULL,
                                          "name" varchar(255) NOT NULL,
    "description" varchar(255) NOT NULL,
    PRIMARY KEY ("id")
    );

CREATE TABLE IF NOT EXISTS "news" (
                                      "id" bigint NOT NULL UNIQUE,
                                      "tag_id" bigint,
                                      "user_id" bigint NOT NULL,
                                      "title" varchar(255) NOT NULL,
    "descriprion" varchar(255) NOT NULL,
    "name_image" varchar(255) NOT NULL,
    PRIMARY KEY ("id")
    );

CREATE TABLE IF NOT EXISTS "news_tags" (
                                           "id_news" bigint NOT NULL,
                                           "id_tag" bigint NOT NULL,
                                           PRIMARY KEY ("id_news", "id_tag")
    );

CREATE TABLE IF NOT EXISTS "tags" (
                                      "id" bigint NOT NULL UNIQUE,
                                      "name" varchar(255) NOT NULL,
    PRIMARY KEY ("id")
    );

ALTER TABLE "comments" ADD CONSTRAINT "comments_fk1" FOREIGN KEY ("user_id") REFERENCES "users"("id");

ALTER TABLE "comments" ADD CONSTRAINT "comments_fk2" FOREIGN KEY ("news_id") REFERENCES "news"("id");

ALTER TABLE "news" ADD CONSTRAINT "news_fk1" FOREIGN KEY ("tag_id") REFERENCES "news_tags"("id_news");

ALTER TABLE "news" ADD CONSTRAINT "news_fk2" FOREIGN KEY ("user_id") REFERENCES "users"("id");

ALTER TABLE "tags" ADD CONSTRAINT "tags_fk0" FOREIGN KEY ("id") REFERENCES "news_tags"("id_tag");