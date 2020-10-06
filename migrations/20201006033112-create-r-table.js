'use strict';

module.exports = {
  up: (queryInterface, Sequelize) => {
    id: {
      allowNull: false,
          autoIncrement: true,
          primaryKey: true,
          type: Sequelize.INTEGER
    },
    name: {
      allowNull: true,
          unique: true,
          type: Sequelize.STRING
    },
    created_at: {
      allowNull: false,
          type: Sequelize.DATE
    },
    updated_at: {
      allowNull: false,
          type: Sequelize.DATE
    }
  },

  down: (queryInterface, Sequelize) => {
      return queryInterface.dropTable("books");
  }
};
