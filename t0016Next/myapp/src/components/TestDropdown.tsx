import React, { useState, useEffect, useMemo  } from 'react';
import Select from 'react-select';
import type { AllData, Vtuber, Movie, KaraokeList } from '../types/singdata';
// import Creatable from 'react-select/creatable' 後々の更新で歌検索として実装したい

// import { colourOptions } from '../data';

// 今は使ってないけど、絶対必要になる。
// const Checkbox = ({ children, ...props }: JSX.IntrinsicElements['input']) => (
//   <label style={{ marginRight: '1em' }}>
//     <input type="checkbox" {...props} />
//     {children}
//   </label>
// );

// DropDinwMo, Kaについは、on~~Seletがnillとか0なら処理を止めべき

const dropStyle =({
     container: (base:any) => ({
      ...base,
      width: '300px'
  }),
  control: (base:any) => ({ //controlでholder、optionで選択肢の文字サイズを指定
    ...base,
    fontSize: '12px'  // ここで文字サイズを指定
    
  }),
}
)

type vtuberOptions = {
  label: string
  value: string
}

 //onVtuberSelectはVtuberが選択されたときに親コンポーネントへ通知するためのコールバック関数
 export const DropDownVt2 = ({posts, onVtuberSelect, onMovieClear, onKaraokeClear }) => {
  const handleVtuberClear = () => {
    onVtuberSelect(null);   //nillになる
    onMovieClear();         //undefinedになる
    onKaraokeClear();
  };
  const [vtuberOptions, setVtuberOptions] = useState<vtuberOptions[]>([]);

  useEffect(() => {
    const fetchVtubers = () => { //適切な名前が思いつけば変える
      try {
        console.log("API Response Vt:", posts.vtubers);
        let havingVt = posts.vtubers.map((vtuber:Vtuber) => ({
          value: vtuber.VtuberId,
          label: vtuber.VtuberName
        }));
        setVtuberOptions(havingVt);
      } catch (error) {
        console.error("Error fetching vtubers:", error);
      }
    };
    fetchVtubers();
  }, []);
  return (
    <>
      <Select
        placeholder="Vtuber名を検索/選択"  //  defaultValue= で何か変わる
        className="basic-single"
        classNamePrefix="select"
        isClearable={true}
        isSearchable={true}
        name="VTuber"  //何のためにあるかわからん
        options={vtuberOptions} //選択候補
        defaultMenuIsOpen={true}
        blurInputOnSelect={true}
        // isDisabled={false} isRtl={false} Value= //その他オプションのメモ
        styles={dropStyle}
        // value={onVtuberSelect}
        onChange={option => { //ここでSelectで選んだものがoptionに格納されるのか？
            // 要：選択中のmovieをクリアする関数
          if (option) {
            console.log("Selected Vtuber value:", option.value);
            onVtuberSelect(option.value);
          } else {
            handleVtuberClear();
          }
        }}
      />
    </>
  );
};

// type DropDownMoProps = {
//   selectedVtuber: Vtuber;
//   onMovieSelect:Movie;
// };
// export const DropDownMo: React.FC<DropDownMoProps> = ({ selectedVtuber ,onMovieSelect}) => {
 
// 
// 
//  movie用
export const DropDownMo2 = ({ posts, selectedVtuber, onMovieSelect, onKaraokeClear }) => {
// export const DropDownMo = ({ selectedVtuber ,onMovieSelect}) => {  
  //const [selectedVtuber, setSelectedVtuber] = useState(null);
  const handleMovieClear = () => {
    onMovieSelect(null);
    onKaraokeClear();
  };
  const [movieOptions, setMovieOptions] = useState([]);
  const [selectedMovie, setSelectedMovie] = useState(null);

  useEffect(() => {
    if (!selectedVtuber) {
      setMovieOptions([]); // Vtuberが選択解除された場合、movieの選択肢を空にする
      return;
    }
    const fetchMovies = async () => {
      try {
        console.log("selectedV=",selectedVtuber)
        console.log("posts.movies=",posts.movies)
        const choicesMovie = posts.movies.filter((movies:Movie) => movies.VtuberId === selectedVtuber);   
        console.log("API Response Mo:", choicesMovie);
        let havingMovie = choicesMovie.map((movie:Movie) => ({
          value: movie.MovieUrl,
          label: movie.MovieTitle
        }));
        setMovieOptions(havingMovie);
      } catch (error) {
        console.error("Error fetching vtubers:", error);
      }
        setSelectedMovie(null);
    };
    fetchMovies();
  }, [selectedVtuber]);
  return (
    <>
      <Select
      placeholder="動画タイトルを検索/選択"
        className="basic-single"
        classNamePrefix="select"
        // value={changeValue}
        isClearable={true}
        isSearchable={true}
        name="movie"
        options={movieOptions}
        blurInputOnSelect={true}  //defaultでtrueなら不要。スマホでアクセスしないと確認できないと思う。
        captureMenuScroll={true} //スマホ用、タブレット用。使ってみてからt/f判断。
        // value={vtuberOptions.find(opt => opt.value === selectedVtuber)}
        // loadingMessage="loading" 
        styles={dropStyle}
        // value={selectedMovie}
        onChange={option => {
          console.log("movieのvalue 前", selectedMovie, "optin=", option?.value)
            // 要：選択中のkaraokeをクリアする関数
          if (option) {
            onMovieSelect(option.value);
            setSelectedMovie(option);
          console.log("movieのvalue 中", selectedMovie)
          } else {
            handleMovieClear();
            setMovieOptions([]);
          }
          console.log("movieのvalue 後", selectedMovie)

        }}
      />
    </>
  );
};


type DropDownKa2Props = {
  posts: AllData;
  selectedMovie: string;
  onKaraokeSelect: number;
};

// karaoke_list用
export const DropDownKa2: React.FC<DropDownKa2Props>  = ({ posts, selectedMovie, onKaraokeSelect }) => {
  // const [movies, setData2] = useState<KaraokeList[]>();
  const [karaokeListOptions, setMovieOptions] = useState([]);
  const [selectedKaraoke, setSelectedKaraoke] = useState<number>(0);

  useEffect(() => {
    if (!selectedMovie) {
      setMovieOptions([]); // Movieが選択解除された場合、karaokeの選択肢を空にする
      return;
    }
    // if (!selectedKaraoke){
    //   setSelectedKaraoke(null);
    //   return;
    // }
    const fetchKaraokes = async () => {
      try {
        console.log("191 selectedMovie=",selectedMovie)
        console.log("karaokes=",posts.karaokes)
        console.log("karaokes.MovieUrl=",posts.karaokes[0].MovieUrl)
        // const movieUrl = {selectedMovie}
        const choiceKaraoke = posts.karaokes.filter((karaokes:KaraokeList) => karaokes.MovieUrl === selectedMovie);
        console.log("API Response ka:", choiceKaraoke );
        let havingkaraokeList = choiceKaraoke.map((karaoke:KaraokeList) => ({
          value: karaoke.KaraokeListId,
          label: karaoke.SongName
        }));
        setMovieOptions(havingkaraokeList);
      } catch (error) {
        console.error("Error fetching vtubers:", error);
      }
      setSelectedKaraoke(0);
    };
    fetchKaraokes();
  }, [selectedMovie]);
  return (
    <>
      <Select
      placeholder="歌を検索/選択"
        className="basic-single"
        classNamePrefix="select"
        isClearable={true}
        isSearchable={true}
        options={karaokeListOptions}
        // value={selectedKaraoke}
        isMulti={true} 
        backspaceRemovesValue={false}
        blurInputOnSelect={true}
        styles={dropStyle}
      onChange={option => {
        // if (option){ //選んだ時にエラー吐く
        console.log("Selected Karaoke1:", option); 
        if (typeof onKaraokeSelect === 'function') {
          onKaraokeSelect(option);
        }    
            // }  
      }}        
      />
    </>
  );
};


    // 公式
    // https://react-select.com/home

    // <select  value={hoge} 
    //これがあると、その値が変化したときのみUIが変化する
    // ない場合は「制御されないコンポーネント」となり、どんな変更でもUIが変化する