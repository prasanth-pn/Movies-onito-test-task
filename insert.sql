\COPY movies (tconst,title_type,primay_title,runtime_minutes,genres) from 'movies.csv' WITH  DELIMITER ',';
\COPY ratings (tconst,average_rating,num_votes) FROM 'ratings.csv' WITH DELIMITER ',';