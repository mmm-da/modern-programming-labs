import React from 'react';
import {Routes, Route} from "react-router-dom";
import MainPage from "./pages/MainPage";
import NoteListPage from "./pages/NoteListPage";

const Router = () => {
    return (
        <Routes>
            <Route path="/" element={<MainPage/>}/>
            <Route path="/notes" element={<NoteListPage/>}/>
        </Routes>
    )
}

export default Router;