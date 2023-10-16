// import { useEffect, useState } from 'react';

// import Link from 'next/link';
// // import YouTube from 'react-youtube';
// import style from '../Youtube.module.css';
// import type { AllData, Vtuber, VtuberMovie } from '../types/singdata'; //type{}で型情報のみインポート
// // import DeleteButton from '../components/DeleteButton';
// import { useRouter } from 'next/router';
// import https from 'https';
// import axios from 'axios';
// import { AxiosRequestConfig } from 'axios';
// import {  Menu, MenuItem,  MenuButton, SubMenu} from '@szhsin/react-menu';
// import '@szhsin/react-menu/dist/index.css';

// import Dropdown from '../components/Dropdown';
// import {DropSelectVt, DropSelectMo} from '../components/Dropselect';

//   type PostsAndCheckSignin= {
//     posts: any; // anyは避けるべき？
//     checkSignin: boolean;
//   }

// //分割代入？
// // 型注釈IndexPage(posts: Post)
// const AllDatePage: React.FC<PostsAndCheckSignin> = ({ posts, checkSignin }) =>  {
//   // data1というステートを定義。streamerの配列を持つ。
//   // setData1はステートを更新する関数。
// const [vtubers, setData1] = useState<Vtuber[]>();
// const [movies, setData2] = useState<VtuberMovie[]>();
// const [selectedVtuber, setSelectedVtuber] = useState(null);
// //   const [check, setData3] = useState<boolean>();
// //   const router = useRouter();

//   useEffect(() => {
//     if (posts) {
//       setData1(posts.vtubers); //上のsetData1の左の変数にposts.vtuberを代入(更新)
//       setData2(posts.vtubers_and_movies);
//       // setData3(checkSingin)
//       console.log("checkSignin=", checkSignin)
//     }
//   }, [posts]);//[]内の要素が変更されるたびにuseEffect内の処理が実行される。

//   return ( 
//     <div>
//        <DropSelectVt onVtuberSelect={setSelectedVtuber}/>
//        <DropSelectMo selectedVtuber={selectedVtuber} />

//        {/* <Dropdown vtubers={vtubers} /> */}

//        {/* <Menu menuButton={<MenuButton>Open menu</MenuButton>}>
//         <MenuItem>New File</MenuItem>
//         <SubMenu label="Open">
//           <MenuItem>mebee</MenuItem>
//           <MenuItem>about</MenuItem>
//           <MenuItem>policy</MenuItem>
//         </SubMenu>
//         <MenuItem>Save</MenuItem>
//       </Menu> */}

// </div>
//   )};

// // オレオレ証明書対応のために、証明書を無視してる。
// // 本番環境では変更が必要。
//   export async function getServerSideProps(context: { req: { headers: { cookie: any; }; }; }) {
//     //リクエストにCookieを入れる。    
//     const httpsAgent = new https.Agent({
//         rejectUnauthorized: false
//     });
//     let resData;
//     try {
//         const response = await axios.get('https://localhost:8080', {
//           // 0019だとnullじゃないとサーバーが起動しない。undefinedはダメだとエラーが出る。
//             httpsAgent: process.env.NODE_ENV === "production" ? undefined : httpsAgent
//         });
//         resData = response.data;
//     } catch (error) {
//         console.log("axios.getでerroe:", error)
//     }

//     // console.log("resData: ", resData)
//     console.log("resData=", resData) //この時点では動画情報持ってる
    
//     // Signinしていればtrueを返す
//     const rawCookie = context.req.headers.cookie;
//     const sessionToken = rawCookie?.split(';').find(cookie => cookie.trim().startsWith('auth-token='))?.split('=')[1];
//     var CheckSignin = false
//     if(sessionToken){CheckSignin = true}

//     console.log("checkSingin=", CheckSignin)

// //    var Posts: PostsAndCheckSignin ;
// //    Posts.checkSignin = CheckSignin ;
// //    Posts. = resData ;

//     return {
//         props: {
//             posts: resData, 
//             checkSignin: CheckSignin
//           }
//     };
// }
  
//   export default AllDatePage;
  


// //   **** memo ****

// //      <DeleteButton Unique_id={item.streamer_id} /> 

// // 各種リンク
//             //  {item2.Movies ? ( <td>{item2.Movies'MovieId'}</td>
//             //   ) : ( <td>未登録</td>      )}

//             //  {item2.MovieUrl ? ( <td>{item2.MovieUrl}</td>
//             //   ) : (<td>-</td>            )}
//             // {item2.MovieTitle ? (  <td>{item2.MovieTitle}</td>
//             //   ) : (<td>-</td>            )} 

// // SSR化する前のコード
// // function AllDatePage( posts : AllData)  {
// //     // data1というステートを定義。streamerの配列を持つ。
// //     // setData1はステートを更新する関数。
// //   const [streamers, setData1] = useState<Streamer[]>();
// //   const [streamerstoMovies, setData2] = useState<StreamerMovie[]>();
// //   const router = useRouter();

// //   useEffect(() => {  
// //     async function fetchData() {
// //         try {
// //             const response = await fetch('https://localhost:8080');
// //                 if (!response.ok) {
// //                     throw new Error('Network response was not ok');
// //                 }
// //             const resData = await response.json();
// //             setData1(resData.streamers);
// //             setData2(resData.streamers_and_moviesmovies);
// //         } catch (error) {
// //             console.error("There was a problem with the fetch operation:", error);
// //         }
// //     }
// //     fetchData();
// // }, []);

// //   return (