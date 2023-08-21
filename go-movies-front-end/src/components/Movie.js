import { useState, useEffect } from 'react'
import { useParams } from 'react-router-dom'

const Movie = () => {
  const [movie, setMovie] = useState({})
  let { id } = useParams();

  useEffect(() => {
    let myMovie = {
      id: 1,
      title: 'Highlander',
      release_date: '1986-03-07',
      runtime: 116,
      mpaa_rating: 'R',
      rating: 7.2,
      poster_url: 'https://image.tmdb.org/t/p/w500/1Q8f9OKNdh5hEX0gyEZHPs0LVcj.jpg',
      description: 'He fought his first battle on the Scottish Highlands in 1536. He will fight his greatest battle on the streets of New York City in 1986. His name is Connor MacLeod. He is immortal.'
    }

    setMovie(myMovie)

  }, [id])

  return (
    <div>
      <h2>Movie: {movie.title}</h2>
      <small><em>{movie.release_date}, {movie.runtime} minutes, Rated {movie.mpaa_rating}</em></small>
      <hr />
    </div>
  )
}

export default Movie