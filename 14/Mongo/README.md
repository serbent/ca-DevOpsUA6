```
docker exec -it  mongodb mongosh -u root -p example 

use movieapp

db.movies.insertMany([
        { title: "The Shawshank Redemption", year: 1994, rating: 8.9, genre: "Drama" },
        { title: "The Godfather", year: 1972, rating: 9.2, genre: "Crime" },
        { title: "The Dark Knight", year: 2008, rating: 9.0, genre: "Action" },
        { title: "Pulp Fiction", year: 1994, rating: 8.9, genre: "Crime" },
        { title: "Forrest Gump", year: 1994, rating: 8.8, genre: "Drama" },
        { title: "Inception", year: 2010, rating: 8.8, genre: "Sci-Fi" },
        { title: "Fight Club", year: 1999, rating: 8.8, genre: "Drama" },
        { title: "The Matrix", year: 1999, rating: 8.7, genre: "Sci-Fi" },
        { title: "Goodfellas", year: 1990, rating: 8.7, genre: "Crime" },
        { title: "Interstellar", year: 2014, rating: 8.6, genre: "Sci-Fi" }
])

db.movies.find().pretty()

db.movies.find({genre: "Sci-Fi"})

db.movies.find({ year: { $gt: 2000} })

db.movies.countDocuments()

db.movies.updateOne(
    { title: "Inception" },
    { $set: {genre: "Science Fiction" } }
)

db.movies.deleteOne({ title: "Interstellar" })

db.movies.insertOne({ title: "The Dark Knight", year: 2008, rating: 9.0, genre: "Action", director: "Torent Quentino"})

db.movies.insertOne({ name:"Me", surname: "you" })
```