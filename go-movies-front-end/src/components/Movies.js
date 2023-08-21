import { useEffect, useState } from 'react'
import { Link } from 'react-router-dom';

const Movies = () => {
  const [movies, setMovies] = useState([])

  useEffect(() => {
    let moviesList = [
      {
        id: 1,
        title: 'Highlander',
        release_date: '1986-03-07',
        runtime: 116,
        mpaa_rating: 'R',
        rating: 7.2,
        poster_url: 'https://image.tmdb.org/t/p/w500/1Q8f9OKNdh5hEX0gyEZHPs0LVcj.jpg',
        description: 'He fought his first battle on the Scottish Highlands in 1536. He will fight his greatest battle on the streets of New York City in 1986. His name is Connor MacLeod. He is immortal.'
      },
      {
        id: 2,
        title: 'Raiders of the Lost Ark',
        release_date: '1981-06-12',
        runtime: 115,
        mpaa_rating: 'PG-13',
        rating: 9,
        poster_url: 'https://image.tmdb.org/t/p/w500/1Q8f9OKNdh5hEX0gyEZHPs0LVcj.jpg',
        description: 'In 1936, American archaeologist Indiana Jones recovers a Golden Idol from a booby-trapped Peruvian temple. Rival archaeologist René Belloq corners him and steals the idol; Jones escapes in a waiting seaplane. After returning to the United States, Jones is briefed by two Army Intelligence agents that Nazi German forces are excavating at Tanis, Egypt, and one of their telegrams mentions Jones former mentor Abner Ravenwood.Jones deduces that the Nazis are seeking the Ark of the Covenant, which Adolf Hitler believes will make their army invincible.The agents recruit Jones to recover the Ark first.'
      }
    ];

    setMovies(moviesList)
  }, [])



  return (
    <div>
      <h2>Movies</h2>
      <hr />

      <table className='table table-striped table-hover'>
        <thead>

          <tr>
            <th>Title</th>
            <th>Release Date</th>
            <th>MPAA Rating</th>
            <th></th>
          </tr>
        </thead>
        <tbody>
          {movies.map(movie => (
            <tr key={movie.id}>
              <td>
                <Link to={`/movies/${movie.id}`}>
                  {movie.title}
                </Link>
              </td>
              <td>{movie.release_date}</td>
              <td>{movie.mpaa_rating}</td>
              <td>
                <a href={`/admin/movie/${movie.id}`} className='btn btn-sm btn-outline-primary'>Edit</a>
              </td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  )
}

export default Movies