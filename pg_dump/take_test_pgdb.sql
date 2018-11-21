--
-- PostgreSQL database dump
--

-- Dumped from database version 10.6 (Ubuntu 10.6-0ubuntu0.18.04.1)
-- Dumped by pg_dump version 10.6 (Ubuntu 10.6-0ubuntu0.18.04.1)

-- Started on 2018-11-26 10:46:27 IST

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

DROP DATABASE take_test_db;
--
-- TOC entry 2957 (class 1262 OID 16385)
-- Name: take_test_db; Type: DATABASE; Schema: -; Owner: postgres
--

CREATE DATABASE take_test_db WITH TEMPLATE = template0 ENCODING = 'UTF8' LC_COLLATE = 'en_IN' LC_CTYPE = 'en_IN';


ALTER DATABASE take_test_db OWNER TO postgres;

\connect take_test_db

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

--
-- TOC entry 1 (class 3079 OID 13039)
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;


--
-- TOC entry 2960 (class 0 OID 0)
-- Dependencies: 1
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';


SET default_tablespace = '';

SET default_with_oids = false;

--
-- TOC entry 196 (class 1259 OID 16386)
-- Name: questions_ans; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.questions_ans (
    id bigint NOT NULL,
    question text,
    option_a text,
    option_b text,
    option_c text,
    option_d text,
    answer character(1),
    created_on date,
    serial_no bigint
);


ALTER TABLE public.questions_ans OWNER TO postgres;

--
-- TOC entry 197 (class 1259 OID 16392)
-- Name: questions_ans_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.questions_ans_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.questions_ans_id_seq OWNER TO postgres;

--
-- TOC entry 2961 (class 0 OID 0)
-- Dependencies: 197
-- Name: questions_ans_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.questions_ans_id_seq OWNED BY public.questions_ans.id;


--
-- TOC entry 201 (class 1259 OID 16466)
-- Name: test_answare_id; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.test_answare_id
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.test_answare_id OWNER TO postgres;

--
-- TOC entry 200 (class 1259 OID 16447)
-- Name: test_answers; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.test_answers (
    id bigint DEFAULT nextval('public.test_answare_id'::regclass) NOT NULL,
    question_id bigint NOT NULL,
    test_id bigint NOT NULL,
    option_selected character(1) NOT NULL
);


ALTER TABLE public.test_answers OWNER TO postgres;

--
-- TOC entry 198 (class 1259 OID 16415)
-- Name: test_details_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.test_details_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.test_details_id_seq OWNER TO postgres;

--
-- TOC entry 199 (class 1259 OID 16417)
-- Name: test_details; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.test_details (
    id bigint DEFAULT nextval('public.test_details_id_seq'::regclass) NOT NULL,
    ref_no character varying(200),
    stud_first_name character varying(100),
    stud_last_name character varying(100),
    stud_mobile_no character varying(100),
    stud_email character varying(100),
    score character varying(10),
    taken_on timestamp(4) with time zone,
    start_at timestamp(4) with time zone,
    ended_at timestamp(4) with time zone
);


ALTER TABLE public.test_details OWNER TO postgres;

--
-- TOC entry 2798 (class 2604 OID 16400)
-- Name: questions_ans id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.questions_ans ALTER COLUMN id SET DEFAULT nextval('public.questions_ans_id_seq'::regclass);


--
-- TOC entry 2946 (class 0 OID 16386)
-- Dependencies: 196
-- Data for Name: questions_ans; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.questions_ans (id, question, option_a, option_b, option_c, option_d, answer, created_on, serial_no) VALUES (6, 'A rectangular park 60 m long and 40 m wide has two concrete crossroads running in the middle of the park and rest of the park has been used as a lawn. If the area of the lawn is 2109 sq. m, then what is the width of the road? ', '2.91 m', '3 m', '5.82 m', 'None of these', 'B', '2018-11-25', 3);
INSERT INTO public.questions_ans (id, question, option_a, option_b, option_c, option_d, answer, created_on, serial_no) VALUES (5, 'A train running at the speed of 60 km/hr crosses a pole in 9 seconds. What is the length of the train?', '120 metres', '180 metres', '324 metres
', '150 metres
', 'D', '2018-11-25', 2);
INSERT INTO public.questions_ans (id, question, option_a, option_b, option_c, option_d, answer, created_on, serial_no) VALUES (7, 'Two trains running in opposite directions cross a man standing on the platform in 27 seconds and 17 seconds respectively and they cross each other in 23 seconds. The ratio of their speeds is:', '1:3
', '3:2
', '3:4', 'None of these', 'B', '2018-11-25', 4);
INSERT INTO public.questions_ans (id, question, option_a, option_b, option_c, option_d, answer, created_on, serial_no) VALUES (8, 'A train travelling at a speed of 75 mph enters a tunnel 31/2 miles long. The train is 1/4 mile long. How long does it take for the train to pass through the tunnel from the moment the front enters to the moment the rear emerges?', '2.5 min', '3 min', '3.2 min', '3.5 min', 'B', '2018-11-25', 5);
INSERT INTO public.questions_ans (id, question, option_a, option_b, option_c, option_d, answer, created_on, serial_no) VALUES (9, 'A watch which gains uniformly is 2 minutes low at noon on Monday and is 4 min. 48 sec fast at 2 p.m. on the following Monday. When was it correct?', '2 p.m. on Tuesday', '2 p.m. on Wednesday', '3 p.m. on Thursday', '1 p.m. on Friday', 'B', '2018-11-25', 6);
INSERT INTO public.questions_ans (id, question, option_a, option_b, option_c, option_d, answer, created_on, serial_no) VALUES (10, 'A boat running upstream takes 8 hours 48 minutes to cover a certain distance, while it takes 4 hours to cover the same distance running downstream. What is the ratio between the speed of the boat and speed of the water current respectively?', '2 : 1', '3 : 2', '8 : 3', 'Cannot be determined', 'C', '2018-11-25', 7);
INSERT INTO public.questions_ans (id, question, option_a, option_b, option_c, option_d, answer, created_on, serial_no) VALUES (11, 'In one hour, a boat goes 11 km/hr along the stream and 5 km/hr against the stream. The speed of the boat in still water (in km/hr) is:', '3 km/hr', '5 km/hr', '8 km/hr', '9 km/hr', 'C', '2018-11-25', 8);
INSERT INTO public.questions_ans (id, question, option_a, option_b, option_c, option_d, answer, created_on, serial_no) VALUES (12, 'This Stock Exchange launched a mobile application and web-based platform for retail investors to buy government securities (G-Secs).', 'Bangalore Stock Exchange', 'National Stock Exchange', 'National Stock Exchange', 'Ahmedabad Stock Exchange', 'B', '2018-11-25', 9);
INSERT INTO public.questions_ans (id, question, option_a, option_b, option_c, option_d, answer, created_on, serial_no) VALUES (13, 'The banker''s discount on Rs. 1600 at 15% per annum is the same as true discount on Rs. 1680 for the same time and at the same rate. The time is:', '3 months', '4 months', '6 months', '8 months', 'B', '2018-11-25', 10);
INSERT INTO public.questions_ans (id, question, option_a, option_b, option_c, option_d, answer, created_on, serial_no) VALUES (14, 'There are two examinations rooms A and B. If 10 students are sent from A to B, then the number of students in each room is the same. If 20 candidates are sent from B to A, then the number of students in A is double the number of students in B. The number of students in room A is:', '20
', '80', '100', '200', 'C', '2018-11-25', 11);
INSERT INTO public.questions_ans (id, question, option_a, option_b, option_c, option_d, answer, created_on, serial_no) VALUES (15, 'The price of 10 chairs is equal to that of 4 tables. The price of 15 chairs and 2 tables together is Rs. 4000. The total price of 12 chairs and 3 tables is:', 'Rs. 3500', 'Rs. 3750', 'Rs. 3840', 'Rs. 3900', 'D', '2018-11-25', 12);
INSERT INTO public.questions_ans (id, question, option_a, option_b, option_c, option_d, answer, created_on, serial_no) VALUES (16, 'The price of 2 sarees and 4 shirts is Rs. 1600. With the same money one can buy 1 saree and 6 shirts. If one wants to buy 12 shirts, how much shall he have to pay ?', 'Rs. 1200', 'Rs. 2400', 'Rs. 4800', 'Cannot be determined', 'B', '2018-11-25', 13);
INSERT INTO public.questions_ans (id, question, option_a, option_b, option_c, option_d, answer, created_on, serial_no) VALUES (17, 'An Informal Gathering occurs when a group of people get together in a casual, relaxed manner. Which situation below is the best example of an Informal Gathering?', 'The book club meets on the first Thursday evening of every month.', 'After finding out about his promotion, Jeremy and a few coworkers decide to go out for a quick drink after work.', 'Mary sends out 25 invitations for the bridal shower she is giving for her sister.', 'Whenever she eats at the Mexican restaurant, Clara seems to run into Peter.', 'B', '2018-11-25', 14);
INSERT INTO public.questions_ans (id, question, option_a, option_b, option_c, option_d, answer, created_on, serial_no) VALUES (18, 'Posthumous Publication occurs when a book is published after the author''s death. Which situation below is the best example of Posthumous Publication ?', 'Posthumous Publication occurs when a book is published after the author''s death. Which situation below is the best example of Posthumous Publication ?', 'Elizabeth is honored with a prestigious literary award for her writing career and her daughter accepts the award on behalf of her deceased mother.', 'Melissa''s publisher cancels her book contract after she fails to deliver the manuscript on time.', 'Clarence never thought he''d live to see the third book in his trilogy published.', 'A', '2018-11-25', 15);
INSERT INTO public.questions_ans (id, question, option_a, option_b, option_c, option_d, answer, created_on, serial_no) VALUES (19, 'Binary number 1101.101 is equivalent to decimal number?', '13.5', '13.75', '13.625', '13.875', 'C', '2018-11-25', 16);
INSERT INTO public.questions_ans (id, question, option_a, option_b, option_c, option_d, answer, created_on, serial_no) VALUES (20, 'A 20 m antenna gives a certain uplink gain at frequencies of 4/6 GHz. For getting same gain in the 20/30 GHz band, antenna size required is metre.', '100', '4', '1', '10', 'B', '2018-11-25', 17);
INSERT INTO public.questions_ans (id, question, option_a, option_b, option_c, option_d, answer, created_on, serial_no) VALUES (22, 'In the context of error detection in satellite transmission, ARQ stands for', 'Automatic Repeat Request', 'Automatic Relay Request', 'Accelerated Recovery Request', 'Automatic Radiation Quenching', 'A', '2018-11-25', 18);
INSERT INTO public.questions_ans (id, question, option_a, option_b, option_c, option_d, answer, created_on, serial_no) VALUES (23, 'The different access methods which permit many satellite users to operate in parallel through a single transponder without interfering with each other as', 'Frequency Division Multiple Access (FDMA)', 'Time Division Multiple Access (TDMA)', 'Code Division Multiple Access (CDMA)', 'All of the above', 'D', '2018-11-25', 19);
INSERT INTO public.questions_ans (id, question, option_a, option_b, option_c, option_d, answer, created_on, serial_no) VALUES (24, 'What is the purpose of Spanning Tree Protocol in a switched LAN?', 'To provide a mechanism for network monitoring in switched environments', 'To prevent routing loops in networks with redundant paths', 'To prevent switching loops in networks with redundant switched paths', 'To manage the VLAN database across multiple switches', 'C', '2018-11-25', 20);
INSERT INTO public.questions_ans (id, question, option_a, option_b, option_c, option_d, answer, created_on, serial_no) VALUES (25, 'You need to create an access list that will prevent hosts in the network range of 192.168.160.0 to 192.168.191.0. Which of the following lists will you use?', 'access-list 10 deny 192.168.160.0 255.255.224.0', 'access-list 10 deny 192.168.160.0 0.0.191.255', 'access-list 10 deny 192.168.160.0 0.0.31.255', 'access-list 10 deny 192.168.0.0 0.0.31.255', 'C', '2018-11-25', 21);
INSERT INTO public.questions_ans (id, question, option_a, option_b, option_c, option_d, answer, created_on, serial_no) VALUES (26, 'Large transformers, when used for some time, become very hot and are cooled by circulating oil. The heating of the transformer is due to', 'the heating effect of current alone', 'hysteresis loss alone', 'both the heating effect of current and hysteresis loss', 'intense sunlight at noon', 'C', '2018-11-25', 22);
INSERT INTO public.questions_ans (id, question, option_a, option_b, option_c, option_d, answer, created_on, serial_no) VALUES (27, 'Radiocarbon is produced in the atmosphere as a result of', 'collision between fast neutrons and nitrogen nuclei present in the atmosphere', 'action of ultraviolet light from the sun on atmospheric oxygen', 'action of solar radiations particularly cosmic rays on carbon dioxide present in the atmosphere', 'lightning discharge in atmosphere', 'A', '2018-11-25', 23);
INSERT INTO public.questions_ans (id, question, option_a, option_b, option_c, option_d, answer, created_on, serial_no) VALUES (28, 'One should not connect a number of electrical appliances to the same power socket because', 'this can damage the appliances due to overloading', 'this can damage the domestic wiring due to overloading', 'this can damage the electrical meter', 'the appliance will not get full voltage', 'B', '2018-11-25', 24);
INSERT INTO public.questions_ans (id, question, option_a, option_b, option_c, option_d, answer, created_on, serial_no) VALUES (29, 'It takes much longer to cook food in the hills than in the plains, because', 'in the hills the atmospheric pressure is lower than that in the plains and therefore water boils at a temperature lower than 100oC causing an increase in cooking time', 'due to low atmospheric pressure on the hills, the water boils at a temperature higher than 100oC and therefore water takes longer to boil', 'in the hills the atmospheric density is low and therefore a lot of heat is lost to the atmosphere', 'in the hills the humidity is high and therefore a lot of heat is absorbed by the atmosphere leaving very little heat for cooking', 'D', '2018-11-25', 25);
INSERT INTO public.questions_ans (id, question, option_a, option_b, option_c, option_d, answer, created_on, serial_no) VALUES (1, 'What is golang', 'functional programming language', 'OO programming language ', 'Both functional and OO  programming language ', 'Non of above', 'A', '2018-11-25', 1);


--
-- TOC entry 2950 (class 0 OID 16447)
-- Dependencies: 200
-- Data for Name: test_answers; Type: TABLE DATA; Schema: public; Owner: postgres
--



--
-- TOC entry 2949 (class 0 OID 16417)
-- Dependencies: 199
-- Data for Name: test_details; Type: TABLE DATA; Schema: public; Owner: postgres
--



--
-- TOC entry 2962 (class 0 OID 0)
-- Dependencies: 197
-- Name: questions_ans_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.questions_ans_id_seq', 30, true);


--
-- TOC entry 2963 (class 0 OID 0)
-- Dependencies: 201
-- Name: test_answare_id; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.test_answare_id', 88, true);


--
-- TOC entry 2964 (class 0 OID 0)
-- Dependencies: 198
-- Name: test_details_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.test_details_id_seq', 64, true);


--
-- TOC entry 2802 (class 2606 OID 16397)
-- Name: questions_ans question_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.questions_ans
    ADD CONSTRAINT question_pkey PRIMARY KEY (id);


--
-- TOC entry 2804 (class 2606 OID 16493)
-- Name: questions_ans serial_no_unique; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.questions_ans
    ADD CONSTRAINT serial_no_unique UNIQUE (serial_no);


--
-- TOC entry 2818 (class 2606 OID 16453)
-- Name: test_answers test_ans_uniqu_id; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.test_answers
    ADD CONSTRAINT test_ans_uniqu_id UNIQUE (id);


--
-- TOC entry 2820 (class 2606 OID 16451)
-- Name: test_answers test_answers_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.test_answers
    ADD CONSTRAINT test_answers_pkey PRIMARY KEY (id);


--
-- TOC entry 2810 (class 2606 OID 16425)
-- Name: test_details test_details_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.test_details
    ADD CONSTRAINT test_details_pkey PRIMARY KEY (id);


--
-- TOC entry 2822 (class 2606 OID 16477)
-- Name: test_answers unique_com_qid_test_id; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.test_answers
    ADD CONSTRAINT unique_com_qid_test_id UNIQUE (test_id, question_id);


--
-- TOC entry 2806 (class 2606 OID 16444)
-- Name: questions_ans unique_id; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.questions_ans
    ADD CONSTRAINT unique_id UNIQUE (id);


--
-- TOC entry 2808 (class 2606 OID 16491)
-- Name: questions_ans unique_question; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.questions_ans
    ADD CONSTRAINT unique_question UNIQUE (question);


--
-- TOC entry 2812 (class 2606 OID 16434)
-- Name: test_details unique_ref_no; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.test_details
    ADD CONSTRAINT unique_ref_no UNIQUE (ref_no);


--
-- TOC entry 2814 (class 2606 OID 16446)
-- Name: test_details unique_test_details_id; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.test_details
    ADD CONSTRAINT unique_test_details_id UNIQUE (id);


--
-- TOC entry 2815 (class 1259 OID 16459)
-- Name: fki_pk_with_question_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX fki_pk_with_question_id ON public.test_answers USING btree (question_id);


--
-- TOC entry 2816 (class 1259 OID 16465)
-- Name: fki_pk_with_test_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX fki_pk_with_test_id ON public.test_answers USING btree (test_id);


--
-- TOC entry 2823 (class 2606 OID 16454)
-- Name: test_answers pk_with_question_id; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.test_answers
    ADD CONSTRAINT pk_with_question_id FOREIGN KEY (question_id) REFERENCES public.questions_ans(id);


--
-- TOC entry 2824 (class 2606 OID 16460)
-- Name: test_answers pk_with_test_id; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.test_answers
    ADD CONSTRAINT pk_with_test_id FOREIGN KEY (test_id) REFERENCES public.test_details(id);


--
-- TOC entry 2959 (class 0 OID 0)
-- Dependencies: 5
-- Name: SCHEMA public; Type: ACL; Schema: -; Owner: postgres
--

GRANT ALL ON SCHEMA public TO PUBLIC;


-- Completed on 2018-11-26 10:46:27 IST

--
-- PostgreSQL database dump complete
--

