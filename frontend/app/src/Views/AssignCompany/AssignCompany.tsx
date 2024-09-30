import React, { useState } from 'react';

const AssignCompany: React.FC = () => {
    const [userId, setUserId] = useState('');
    const [company, setCompany] = useState('');

    const assignCompany = async () => {
        try {
            const res = await fetch('/assign', {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify({ id: userId, company: company }),
            });

            const result = await res.json();
            alert(result.message);
        } catch (error) {
            console.error("Failed to assign company", error);
            alert("Error assigning company");
        }
    };

    return (
        <div>
            <input
                type="text"
                id="userId"
                placeholder="User ID"
                value={userId}
                onChange={(e) => setUserId(e.target.value)}
            />
            <input
                type="text"
                id="company"
                placeholder="Company"
                value={company}
                onChange={(e) => setCompany(e.target.value)}
            />
            <button onClick={assignCompany}>Assign</button>
        </div>
    );
};

export default AssignCompany;
