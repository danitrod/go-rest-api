create table personalities(
    id serial primary key,
    name varchar,
    story varchar
);

INSERT INTO personalities(name, story) VALUES
('Mahatma Gandhi', 'Mohandas Karamchand Gandhi (2 October 1869 – 30 January 1948) was an Indian lawyer, anti-colonial nationalist and political ethicist who employed nonviolent resistance to lead the successful campaign for India''s independence from British rule, and to later inspire movements for civil rights and freedom across the world. The honorific Mahātmā (Sanskrit: "great-souled", "venerable"), first applied to him in 1914 in South Africa, is now used throughout the world.'),
('Martin Luther King Jr.', 'Martin Luther King Jr. (born Michael King Jr.; January 15, 1929 – April 4, 1968) was an American Baptist minister and activist who became the most visible spokesman and leader in the civil rights movement from 1955 until his assassination in 1968. An African American church leader and the son of early civil rights activist and minister Martin Luther King Sr., King advanced civil rights for people of color in the United States through nonviolence and civil disobedience. Inspired by his Christian beliefs and the nonviolent activism of Mahatma Gandhi, he led targeted, nonviolent resistance against Jim Crow laws and other forms of discrimination.');
