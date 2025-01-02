// script.js
document.addEventListener("DOMContentLoaded", () => {
    const expenseForm = document.getElementById("expense-form");
    const expenseTableBody = document.getElementById("expense-table-body");
    const categorySummary = document.getElementById("category-summary");
    const expenseChart = document.getElementById("expense-chart").getContext("2d");

    let expenses = JSON.parse(localStorage.getItem("expenses")) || [];

    const updateUI = () => {
        expenseTableBody.innerHTML = "";
        const categoryTotals = {};

        expenses.forEach((expense, index) => {
            const row = document.createElement("tr");

            row.innerHTML = `
                <td>${expense.amount}</td>
                <td>${expense.description}</td>
                <td>${expense.category}</td>
                <td><button onclick="deleteExpense(${index})">Delete</button></td>
            `;

            expenseTableBody.appendChild(row);

            categoryTotals[expense.category] = (categoryTotals[expense.category] || 0) + parseFloat(expense.amount);
        });

        categorySummary.innerHTML = "";
        for (let category in categoryTotals) {
            const li = document.createElement("li");
            li.textContent = `${category}: $${categoryTotals[category].toFixed(2)}`;
            categorySummary.appendChild(li);
        }

        renderChart(categoryTotals);
    };

    const renderChart = (data) => {
        const labels = Object.keys(data);
        const values = Object.values(data);

        new Chart(expenseChart, {
            type: "pie",
            data: {
                labels,
                datasets: [
                    {
                        label: "Spending by Category",
                        data: values,
                        backgroundColor: ["#ff6384", "#36a2eb", "#cc65fe"],
                    },
                ],
            },
        });
    };

    const deleteExpense = (index) => {
        expenses.splice(index, 1);
        localStorage.setItem("expenses", JSON.stringify(expenses));
        updateUI();
    };

    expenseForm.addEventListener("submit", (e) => {
        e.preventDefault();

        const amount = document.getElementById("amount").value;
        const description = document.getElementById("description").value;
        const category = document.getElementById("category").value;

        expenses.push({ amount, description, category });
        localStorage.setItem("expenses", JSON.stringify(expenses));
        updateUI();

        expenseForm.reset();
    });

    window.deleteExpense = deleteExpense;

    updateUI();
});
