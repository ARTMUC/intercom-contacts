import React from 'react';
import {Contact} from "../../Models/Contact";
import styles from './ContactList.module.css';

interface ContactListProps {
    contacts: Contact[];
}

const ContactList: React.FC<ContactListProps> = ({contacts}) => {
    return (
        <div>
            <table className={styles.table}>
                <thead>
                <tr>
                    <th>ID</th>
                    <th>Name</th>
                    <th>Companies</th>
                </tr>
                </thead>
                <tbody>
                {contacts.map((contact) => (
                    <tr key={contact.id}>
                        <td>{contact.id}</td>
                        <td>{contact.email}</td>
                        <td>{JSON.stringify(contact.companies)}</td>
                    </tr>
                ))}
                </tbody>
            </table>
        </div>
    );
};

export default ContactList;
