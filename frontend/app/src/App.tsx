import React, {useEffect, useState} from 'react';
import ContactList from './Views/ContactList/ContactList';
import './App.css';
import {Pagination} from "./Models/Pagination";
import {Contact} from "./Models/Contact";
import {ApiError} from "./Models/ApiError";

const App: React.FC = () => {
    const [isFetched, setIsFetched] = useState<boolean>(false);
    const [contacts, setContacts] = useState<Contact[]>([]);
    const [error, setError] = useState<string | null>(null);

    useEffect(() => {
        fetchUsers().then()
        // @ts-ignore
        // window.ListContacts()
        //     .then((contacts: any) => setContacts(contacts.map((c: { ID: any; Email: any; Companies: any; }) => {
        //         return {
        //             id: c.ID,
        //             email: c.Email,
        //             companies: c.Companies
        //         }
        //     }) || []))
        //     .catch((e: any) => setError(e))
        //     .finally(() => setIsFetched(true))
    }, [])

    const fetchUsers = async () => {
        try {
            let res = await fetch('/contact');
            if (!res.ok) {
                const errorData = (await res.json()) as ApiError;
                setError(errorData.message)
                return;
            }

            const response = (await res.json()) as Pagination<Contact>;
            setContacts(response.rows || []);
        } catch (error) {
            setError((error as Error).message)
        } finally {
            setIsFetched(true)
        }
    };

    return (
        <div className="App">
            <h1>Intercom contacts</h1>
            {error && <div>{error}</div>}
            {isFetched && contacts?.length === 0 ? <p>No contacts found</p> : ''}
            {isFetched && contacts?.length > 0 ? <ContactList contacts={contacts}/> : ''}
        </div>
    );
};

export default App;
